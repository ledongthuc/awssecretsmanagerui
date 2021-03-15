package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/ledongthuc/awssecretsmanagerui/server/routes"
)

//go:embed static/*
var staticResources embed.FS

type arguments struct {
	Host     string
	Port     int
	Headless bool
}

func main() {
	serverAddr := composeServerAddr()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORS())
	e.HideBanner = true
	routes.SetupRoutes(e, staticResources)

	if err := e.Start(serverAddr); err != nil {
		panic(err)
	}
}

func composeServerAddr() string {
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3000"
	}
	return fmt.Sprintf("%s:%s", host, port)
}
