package server

import (
	"context"
	"database/sql"
	"net/http"
	"yomikyasu/internal/database"
	"yomikyasu/internal/dto"
	"yomikyasu/internal/model"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterFeedRoutes(e *echo.Echo) {
	e.GET("/feeds", listFeeds(s.db))
	e.POST("/feeds", createFeed(s.db))
}

func createFeed(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		feed := &dto.Feed{}
		if err := c.Bind(feed); err != nil {
			return err
		}

		result, err := db.Query().CreateFeed(context.Background(), model.CreateFeedParams{
			Url:        feed.Url,
			IsFullText: feed.IsFullText,
			ItemSince: sql.NullFloat64{
				Float64: feed.ItemSince,
			},
			MaxItems: sql.NullInt64{
				Int64: feed.MaxItems,
			},
			Language: sql.NullString{
				String: feed.Language,
			},
		})

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, result)
	}
}

func listFeeds(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		feeds, err := db.Query().ListFeeds(context.Background())

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, feeds)
	}
}
