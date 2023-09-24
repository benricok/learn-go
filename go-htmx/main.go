package main

import (
	"go-htmx/database"
	"log"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		db_url = "./app.db"
	}

	_, err := database.NewDatabase(db_url)
	if err != nil {
		log.Fatalf("Could not init db: %+v", err)
	}

	tmpl, err := template.ParseFiles(
		"./public/index.html",
	)

	if err != nil {
		log.Fatalf("Could not initialise templates: %+v", err)
	}

    e := echo.New()
	e.Renderer = endpoints.NewTemplateRenderer(tmpl)
	e.Use(middleware.Logger())

	e.GET("/", endpoints.HandleIndex)

    e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}