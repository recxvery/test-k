package main

import (
	"context"
	"kaspersky-testovoe/internal/db"
	"kaspersky-testovoe/internal/service"
	"kaspersky-testovoe/internal/transport"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("Ошибка подгружения локальных переменных + ", err)
	}

	initCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	new_pool, err := db.ConnectDB(initCtx)
	if err != nil {
		log.Println("Can't connect to DB")
		panic(err)
	}

	defer new_pool.Close()

	m, err := migrate.New(
		"file://migrations",
		os.Getenv("CONN_STRING"),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != migrate.ErrNoChange {
		log.Println("Ошибка миграции ", err)
		panic(err)
	}

	newRepository := db.NewRepostoryDB(new_pool)
	service := service.NewThreatService(newRepository)

	transport.StartServer(service)
}
