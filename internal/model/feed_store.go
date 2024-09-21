package model

import (
	"database/sql"
	"log"
	"reflect"
	"yomikyasu/internal/database"
)

const (
	CREATE_FEED = "INSERT INTO feeds (url, is_full_text, item_since, max_items, language) VALUES (?, ?, ?, ?, ?)"
	GET_FEEDS   = "SELECT * FROM feeds"
	DELETE_FEED = "DELETE FROM feeds where id = ?"
)

func CreateFeed(db *database.Service, feed *Feed) (*sql.Result, error) {
	stmt, _ := (*db).Prepare(CREATE_FEED)
	defer stmt.Close()

	result, err := stmt.Exec(feed.Url, feed.IsFullText, feed.ItemSince, feed.MaxItems, feed.Language)

	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	return &result, nil
}

func GetFeeds(db *database.Service) ([]Feed, error) {
	stmt, _ := (*db).Prepare(GET_FEEDS)
	defer stmt.Close()

	result, err := (*stmt).Query()

	if err != nil {
		return nil, err
	}

	feeds := make([]Feed, 0)
	for result.Next() {
		c := Feed{}

		s := reflect.ValueOf(&c).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := result.Scan(columns...)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(c)
		feeds = append(feeds, c)
	}

	return feeds, nil
}
