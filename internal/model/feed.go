package model

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
