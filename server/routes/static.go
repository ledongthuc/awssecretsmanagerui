package routes

import (
	"os"

	"github.com/labstack/echo"
)

func setupStaticResourceRoutes(g *echo.Group) {
	g.Static("/", getStaticPath())
}

func getStaticPath() string {
	asset, exist := os.LookupEnv("STATIC_PATH")
	if !exist {
		asset = "static/"
	}
	return asset
}
