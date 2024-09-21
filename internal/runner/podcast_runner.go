package runner

import (
	"context"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
	"yomikyasu/internal/database"
	"yomikyasu/internal/model"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/mmcdole/gofeed"
	"golang.org/x/sync/errgroup"
)

type WorkerRequest struct {
	// Item for this request
	Item *gofeed.Item

	// Directory to which the file wil write to
	Directory string

	// Language of the item
	LanguageCode string

	// Whether to use natural Voice
	UseNaturalVoice bool

	// Speed of Synthesized Speech
	SpeechSpeed float64
}

type Runner interface {
	Run(ctx context.Context)
	RunOnce(ctx context.Context)
}

type runner struct {
	wg      *sync.WaitGroup
	db      *database.Service
	fp      *gofeed.Parser
	channel chan *WorkerRequest
}

const SPEECH_SYNTHESIZE_RETRY_CNT = 5
const FEED_RETRY_CNT = 5

var (
	refreshInterval, _ = strconv.Atoi(os.Getenv("REFERSH_INTERVAL"))
	workerSize, _      = strconv.Atoi(os.Getenv("WORKER_SIZE"))
	ttsClient          *texttospeech.Client
	useNaturalVoice    = true
	speechSpeed        = 1.25
)

func New(ctx context.Context, dbService database.Service) Runner {
	if ttsClient != nil {
		ttsClient, _ = texttospeech.NewClient(ctx)
	}

	var wg sync.WaitGroup
	wg.Add(workerSize)

	channel := make(chan *WorkerRequest, 100)
	for i := 0; i < workerSize; i++ {
		go processSpeechGeneration(&wg, ttsClient, channel, ctx)
	}

	return &runner{
		db:      &dbService,
		wg:      &wg,
		fp:      gofeed.NewParser(),
		channel: channel,
	}
}

func processSpeechGeneration(wg *sync.WaitGroup, client *texttospeech.Client, workerItems chan *WorkerRequest, ctx context.Context) error {
	defer wg.Done()

	for workerItem := range workerItems {
		feedItem := workerItem.Item

		log.Printf("Start procesing %v ", feedItem.Title)

		speechRequests := rss.GetSynthesizeSpeechRequests(feedItem, workerItem.LanguageCode, workerItem.UseNaturalVoice, workerItem.SpeechSpeed)
		audioContent := make([]byte, 0)

		for _, ssr := range speechRequests {
			var err error = nil
			var resp *texttospeechpb.SynthesizeSpeechResponse = nil
			for i := 0; i < SPEECH_SYNTHESIZE_RETRY_CNT; i++ {
				if i > 0 {
					log.Printf("Retry speech synthesize in 1 second due to error %v, count: %v", err, i)
					time.Sleep(time.Second)
				}

				resp, err = client.SynthesizeSpeech(ctx, ssr)
				if err != nil {
					log.Printf("Error Encountered, Response: %v\n", err.Error())
					continue
				}

				if len(resp.AudioContent) > 0 {
					audioContent = append(audioContent, resp.AudioContent...)
					break
				}
			}
			if err != nil {
				return err
			}
		}

		log.Printf("Finished Processing: %v, written to %v\n", feedItem.Title)
	}

	return nil
}

// get all feeds
// run text to speech with worker
// insert into db.
func (r *runner) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(time.Hour.Hours() * float64(refreshInterval)))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			r.RunOnce(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (r *runner) RunOnce(ctx context.Context) {
	feeds, _ := model.GetFeeds(r.db)
	g := new(errgroup.Group)

	for feed := range slices.Values(feeds) {
		v := feed
		g.Go(func() error {
			log.Printf("feed: %v\n", v)
			parsedFeed := getFeedWithRetry(r.fp, v.Url)
			processedItems := 0

			if parsedFeed == nil {
				log.Printf("Fail to fetch/parse feed: %v \n", v.Url)
				return nil
			}

			feedLanguage := func(lang string) string {
				if strings.Contains(strings.ToLower(lang), "zh") {
					return "cmn-CN"
				}

				return lang
			}(feed.Language)

			for item := range slices.Values(parsedFeed.Items) {
				if processedItems == v.MaxItems {
					return nil
				}
				if time.Since(item.PublishedParsed.Local()).Hours() <= v.ItemSince {
					log.Printf("Adding item... title: %s", item.Title)
					r.channel <- &WorkerRequest{
						Item:            item,
						LanguageCode:    feedLanguage,
						UseNaturalVoice: useNaturalVoice,
						SpeechSpeed:     speechSpeed,
					}

				}
			}

			return nil
		})
	}
}

func getFeedWithRetry(fp *gofeed.Parser, v string) *gofeed.Feed {
	var feed *gofeed.Feed = nil
	var err error = nil

	for i := 0; i < FEED_RETRY_CNT; i++ {
		if i > 0 {
			log.Printf("Retry due to Error GET: %v. \n", err)
			time.Sleep(2000)
		}

		feed, err = fp.ParseURL(v)
		if err == nil && feed != nil {
			return feed
		}
	}

	return feed
}
