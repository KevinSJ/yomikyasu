// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package model

import (
	"database/sql"
)

type Config struct {
	ID                 int64
	UseNaturalVoice    bool
	SpeechSpeed        float64
	FullTextServiceUrl sql.NullString
	RefreshInterval    sql.NullInt64
}

type Episode struct {
	ID           int64
	Uuid         string
	FeedID       int64
	Url          string
	Title        string
	Description  sql.NullString
	PubDate      sql.NullString
	FileSize     sql.NullFloat64
	Duration     sql.NullFloat64
	AudioContent []byte
}

type Feed struct {
	ID         int64
	Url        string
	IsFullText bool
	ItemSince  sql.NullFloat64
	MaxItems   sql.NullInt64
	Language   sql.NullString
}

type Podcast struct {
	ID          int64
	Link        string
	Title       string
	Description sql.NullString
}

type PodcastFeed struct {
	PodcastID int64
	FeedID    int64
}

type Runner struct {
	ID              int64
	WorkerSize      int64
	RefreshInterval int64
}
