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
var DefaultPort = "30301"

type params struct {
	Host string
	Port string
}

func main() {
	p := parseParams()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORS())
	e.HideBanner = true
	if os.Getenv("AUTH_ENABLED") == "true" {
		e.Use(middleware.BasicAuth(routes.Auth))
	}
	routes.SetupRoutes(e, staticResources)

	serverAddr := fmt.Sprintf("%s:%s", p.Host, p.Port)
	if err := e.Start(serverAddr); err != nil {
		panic(err)
	}
}

func parseParams() params {
	p := params{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	if p.Host == "" {
		p.Host = "localhost"
	}
	if p.Port == "" {
		p.Port = DefaultPort
	}
	return p
}
