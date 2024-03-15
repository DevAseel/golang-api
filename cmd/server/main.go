package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/DevAseel/golang-api/db"
	"github.com/DevAseel/golang-api/router"
	"github.com/DevAseel/golang-api/services"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Serve() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	fmt.Println("API is listening on port", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
		Models: services.New(dbConn.DB),
	}

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
