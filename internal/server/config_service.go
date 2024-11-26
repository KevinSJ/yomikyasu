package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"yomikyasu/internal/database"
	"yomikyasu/internal/dto"
	"yomikyasu/internal/model"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterConfigRoutes(e *echo.Echo) {
	e.GET("/configs", listConfigs(s.db))
	e.POST("/configs", createConfig(s.db))
}

func createConfig(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		configParams := &dto.Config{}
		if err := c.Bind(configParams); err != nil {
			return err
		}

		// TODO:  <27-11-24, kevin> Implement validation //
		//c.Validate()

		result, err := db.Query().CreateConfig(context.Background(), model.CreateConfigParams{
			UseNaturalVoice: configParams.UseNaturalVoice,
			SpeechSpeed:     configParams.SpeechSpeed,
			FullTextServiceUrl: sql.NullString{
				String: configParams.FullTextServiceUrl,
			},
		})

		fmt.Printf("result: %v\n", result)
		fmt.Printf("err: %v\n", err)

		return c.JSON(http.StatusCreated, c)
	}
}

func listConfigs(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		configs, err := db.Query().ListConfigs(context.Background())

		if err != nil {
			return err
		}

		fmt.Printf("configs: %v\n", configs)

		return c.JSON(http.StatusOK, configs)
	}
}
