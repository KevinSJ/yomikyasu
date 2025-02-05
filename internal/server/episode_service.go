package server

import (
	"context"
	"log"
	"net/http"
	"yomikyasu/internal/database"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterEpisodeRoutes(e *echo.Echo) {
	e.GET("/episodes/:episodeId/content", getPodcastEpisode(s.db))
}

func getPodcastEpisode(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		episodeId := c.Param("episodeId")
		log.Printf("episodeId: %s", episodeId)
		result, err := db.Query().GetEpisodeContentByUuid(context.Background(), episodeId)

		if err != nil {
			return err
		}

		return c.Blob(http.StatusOK, http.DetectContentType(result), result)
	}
}
