package tool

import (
	"bytes"
	"html/template"
	"yomikyasu/internal/model"
)

const (
	XML_TEMPLATE string = `
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
    <channel>
        <atom:link href="{{.PodLink}}" rel="self" type="application/rss+xml" />
        <title>{{.PodTitle}}</title>
        <description>{{.PodDescription}}</description>
        <link>{{.PodLink}}</link>
        {{range .PodEpisodes}}
        <item>
            <guid>{{.Url}}</guid>
            <link>{{.Url}}</link>
            <title>{{.Title}}</title>
            <description>{{.Description}}</description>
            <pubDate>{{.PubDate}}</pubDate>
            <enclosure url="{{.Url}}" type="audio/mpeg" length="{{.FileSize}}"/>
            <itunes:duration>{{.Duration}}</itunes:duration>
        </item>
        {{end}}
    </channel>
</rss>
    `
)

type Episode struct {
	Url, Title, Description, PubDate string
	FileSize                         int64
	Duration                         float64
}

type Podcast struct {
	PodLink, PodTitle, PodDescription string
	PodEpisodes                       []Episode
}

func GeneratePodcastsXmlFeed(podcastsEpisode *[]model.PodcastsEpisode) (*bytes.Buffer, error) {
	t := template.Must(template.New("xml_feed").Parse(XML_TEMPLATE))

	podcast := Podcast{
		PodLink:        (*podcastsEpisode)[0].PodcastLink,
		PodTitle:       (*podcastsEpisode)[0].PodcastTitle,
		PodDescription: (*podcastsEpisode)[0].Description.String,
	}

	podcast.PodEpisodes = Map(*podcastsEpisode, func(podcastsEpisode model.PodcastsEpisode) Episode {
		return Episode{
			Url:         podcast.PodLink + "/episodes/" + podcastsEpisode.Uuid + "/content",
			Title:       podcastsEpisode.Title,
			Description: podcastsEpisode.Description.String,
			PubDate:     podcastsEpisode.PubDate.String,
			FileSize:    int64(podcastsEpisode.FileSize.Float64),
			Duration:    podcastsEpisode.Duration.Float64,
		}
	})

	buf := new(bytes.Buffer)

	t.Execute(buf, podcast)

	return buf, nil
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
