package server

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"yomikyasu/internal/database"
	"yomikyasu/internal/dto"
	"yomikyasu/internal/model"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterPodcastRoutes(e *echo.Echo) {
	e.GET("/podcasts", listPodcasts(s.db))
	e.POST("/podcasts", createPodcast(s.db))
	e.POST("/podcasts/:podcastId/feeds", addFeedToPodcast(s.db))
	e.GET("/podcasts/:podcastId", getGeneratedPodcastFeed(s.db))
}

func getGeneratedPodcastFeed(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		podcastId := c.Param("podcastId")
		podcastIdInt, _ := strconv.ParseInt(podcastId, 10, 64)
		result, err := db.Query().GetPodcastEpisodesByPodcastId(context.Background(), podcastIdInt)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, result)
	}
}

func addFeedToPodcast(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		podcastFeed := &dto.PodcastFeed{}
		if err := c.Bind(podcastFeed); err != nil {
			return err
		}
		result, err := db.Query().CreatePodcastFeed(context.Background(), model.CreatePodcastFeedParams{
			PodcastID: podcastFeed.PodcastId,
			FeedID:    podcastFeed.FeedId,
		})

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, result)
	}
}

func createPodcast(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		podcast := &dto.Podcast{}
		if err := c.Bind(podcast); err != nil {
			return err
		}

		result, err := db.Query().CreatePodcast(context.Background(), model.CreatePodcastParams{
			Link:        podcast.Link,
			Title:       podcast.Title,
			Description: sql.NullString{},
		})

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, result)
	}
}

func listPodcasts(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		feeds, err := db.Query().ListPodcasts(context.Background())

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, feeds)
	}
}
