package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ledongthuc/awssecretsmanagerui/server/routes"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.CORS())
	routes.SetupRoutes(e)
	e.Logger.Info(e.Start(composeServerAddr()))
}

func composeServerAddr() string {
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	return fmt.Sprintf("%s:%s", host, port)
}
