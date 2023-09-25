package main

import (
	"go-htmx/auth"
	"go-htmx/database"
	"go-htmx/endpoints"
	"html/template"
	"log"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
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
		"./public/login.html",
		"./public/header.html",
		"./public/nav.html",
		"./public/home.html",
		"./public/help.html",
		"./public/settings.html",
	)

	if err != nil {
		log.Fatalf("Could not initialise templates: %+v", err)
	}

    e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = endpoints.NewTemplateRenderer(tmpl)

	e.GET("/", endpoints.HandleIndex)
	e.GET("/css/style.css", func(c echo.Context) error { return c.File("./public/css/style.css")})
	e.GET("/login", endpoints.HandleLoginForm)
	e.POST("/login", endpoints.Login)
	e.GET("/logout", endpoints.Logout)

	app := e.Group("/app")
	{
		app.Use(echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc:	auth.Claim,
			SigningKey: 	[]byte(auth.GetJWTSecret()),
			TokenLookup: 	"cookie:access-token",
			ErrorHandler:	auth.JWTErrorChecker,
		}))

		app.GET("/home", endpoints.HandleHome)
		app.GET("/settings", endpoints.HandleSettings)
		app.GET("/help", endpoints.HandleHelp)
	}

    e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}