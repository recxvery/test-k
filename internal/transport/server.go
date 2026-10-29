package transport

import (

	// "github.com/labstack/echo/middleware"

	"kaspersky-testovoe/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(service *service.ThreatService) {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	threatHandler := NewThreatHandler(service)

	e.GET("/info/:md5", threatHandler.GetThreatHandler)
	e.POST("/info", threatHandler.InsertThreatHandler)
	e.PUT("/info/:md5", threatHandler.UpdateThreatHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
