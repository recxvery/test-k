package http

import (
	"context"
	"kaspersky-testovoe/db"
	"log"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func startServer() {
	e := echo.New()

	e.Use(middleware)
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	ctx := context.Background()
	new_pool, err := db.ConnectDB(ctx)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	d := db.NewRepostoryDB(new_pool)

	threatHandler := NewThreatHandler(d)

	e.GET("/info/:md5", threatHandler.GetThreatHandler)
	e.POST("/info", threatHandler.InsertThreatHandler)
	e.PUT("/info/:md5", threatHandler.UpdateThreatHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
