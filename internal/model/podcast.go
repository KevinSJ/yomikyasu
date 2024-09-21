package model

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
