// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: episodes.sql

package model

import (
	"context"
	"database/sql"
)

const createEpisode = `-- name: CreateEpisode :one
INSERT INTO episodes (
    uuid,
    feed_id,
    url,
    title,
    description,
    pub_date,
    file_size,
    duration,
    audio_content
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, uuid, feed_id, url, title, description, pub_date, file_size, duration, audio_content
`

type CreateEpisodeParams struct {
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

func (q *Queries) CreateEpisode(ctx context.Context, arg CreateEpisodeParams) (Episode, error) {
	row := q.db.QueryRowContext(ctx, createEpisode,
		arg.Uuid,
		arg.FeedID,
		arg.Url,
		arg.Title,
		arg.Description,
		arg.PubDate,
		arg.FileSize,
		arg.Duration,
		arg.AudioContent,
	)
	var i Episode
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.FeedID,
		&i.Url,
		&i.Title,
		&i.Description,
		&i.PubDate,
		&i.FileSize,
		&i.Duration,
		&i.AudioContent,
	)
	return i, err
}

const getEpisodeContentByUuid = `-- name: GetEpisodeContentByUuid :one
SELECT audio_content from episodes where uuid = ?
`

func (q *Queries) GetEpisodeContentByUuid(ctx context.Context, uuid string) ([]byte, error) {
	row := q.db.QueryRowContext(ctx, getEpisodeContentByUuid, uuid)
	var audio_content []byte
	err := row.Scan(&audio_content)
	return audio_content, err
}

const getEpisodeExistsByUrlAndFeedId = `-- name: GetEpisodeExistsByUrlAndFeedId :one
SELECT EXISTS (SELECT 1 FROM episodes WHERE url = ? and feed_id = ? LIMIT 1)
`

type GetEpisodeExistsByUrlAndFeedIdParams struct {
	Url    string
	FeedID int64
}

func (q *Queries) GetEpisodeExistsByUrlAndFeedId(ctx context.Context, arg GetEpisodeExistsByUrlAndFeedIdParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getEpisodeExistsByUrlAndFeedId, arg.Url, arg.FeedID)
	var column_1 int64
	err := row.Scan(&column_1)
	return column_1, err
}
