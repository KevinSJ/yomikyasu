package server

import (
	"context"
	"net/http"
	"yomikyasu/internal/database"
	"yomikyasu/internal/model"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterFeedRoutes(e *echo.Echo) {
	e.GET("/feeds", getAllFeeds(s.db))
	e.POST("/feeds", createFeed(s.db))
}

func createFeed(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		feed := &model.CreateFeedParams{}
		if err := c.Bind(feed); err != nil {
			return err
		}

		result, err := db.Query().CreateFeed(context.Background(), *feed)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, result)
	}
}

func getAllFeeds(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		feeds, err := db.Query().ListFeeds(context.Background())

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, feeds)
	}
}
