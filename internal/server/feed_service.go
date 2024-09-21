package server

import (
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
		feed := &model.Feed{}
		if err := c.Bind(feed); err != nil {
			return err
		}

		result, err := model.CreateFeed(&db, feed)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, result)
	}
}

func getAllFeeds(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		feeds, err := model.GetFeeds(&db)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, feeds)
	}
}
