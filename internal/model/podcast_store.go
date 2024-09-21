package model

import (
	"database/sql"
	"log"
	"yomikyasu/internal/database"
)

const (
	CREATE_EPISODE = "INSERT INTO episodes (uuid, feed_id, url, title, description, pub_date, file_size, duration, audio_content) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

func CreateEpisode(db *database.Service, episode *Episode) (*sql.Result, error) {
	stmt, _ := (*db).Prepare(CREATE_EPISODE)
	defer stmt.Close()

	result, err := stmt.Exec(episode.UUID, episode.FeedId, episode.Url, episode.Title, episode.Description, episode.PubDate, episode.FileSize, episode.Duration, episode.AudioContent)

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return &result, nil
}
