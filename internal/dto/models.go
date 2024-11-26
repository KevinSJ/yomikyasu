package dto


type (
	Config struct {
        Id int
		// UseNaturalVoice determines whether to use natural voice from Google, this only have quota of 1
		// Million, whereas the quota for standard voice is 4 Million
		UseNaturalVoice bool `json:"useNaturalVoice"`
		// Speed of synthesized speech
		SpeechSpeed float64 `json:"speechSpeed"`
		// FullTextServiceUrl a full text service to be used when a feed is not
		//in full text
		FullTextServiceUrl string `json:"fullTextServiceUrl,omitempty"`
	}
)

type (
	Feed struct {
		Id int `json:",omitempty"`
		// Url for the feed
		Url string `json:"url"`

		// IsFullText whether the feed is a full text feed.
		IsFullText bool `json:"isFullText"`

		// ItemSince filter the feed based on the pub time in hours
		ItemSince float64 `json:"itemSince"`

		// MaxItems is the max number of item per feed after applying the time
		// filter
		MaxItems int `json:"maxItems"`

		// Language the language for this rss feed.
		Language string `json:"language"`
	}
)

type (
	Episode struct {
		UUID                             string
		Url, Title, Description, PubDate string
		FileSize                         float64
		Duration                         float64
		FeedId                           int64
		AudioContent                     []byte
	}

	Podcast struct {
		Link, Title, Description string
		Episodes                 []Episode
	}
)

type Runner struct {
	ID              int64
	WorkerSize      int64
	RefreshInterval int64
}
