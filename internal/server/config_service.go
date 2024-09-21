package server

import (
	"fmt"
	"net/http"
	"yomikyasu/internal/database"
	"yomikyasu/internal/model"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterConfigRoutes(e *echo.Echo) {
	e.GET("/configs", getAllConfigs(s.db))
	e.POST("/configs", createConfig(s.db))
}

func createConfig(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		config := &model.Config{}
		if err := c.Bind(config); err != nil {
			return err
		}

		result, err := model.CreateConfig(&db, config)

		fmt.Printf("result: %v\n", result)
		fmt.Printf("err: %v\n", err)

		return c.JSON(http.StatusCreated, c)
	}
}

func getAllConfigs(db database.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		configs, err := model.GetConfigs(&db)

		if err != nil {
			return err
		}

		fmt.Printf("configs: %v\n", configs)

		return c.JSON(http.StatusOK, configs)
	}
}
