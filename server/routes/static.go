package routes

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setupStaticResourceRoutes(g *echo.Group, staticResources embed.FS) {
	var contentHandler = echo.WrapHandler(http.FileServer(http.FS(staticResources)))
	var contentRewrite = middleware.Rewrite(map[string]string{
		"/*": "/static/$1",
	})
	g.GET("", contentHandler, contentRewrite)
	g.GET("/", contentHandler, contentRewrite)
	g.GET("/*", contentHandler, contentRewrite)
	g.GET("/_next/*", contentHandler, contentRewrite)
}
