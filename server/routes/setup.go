package routes

import (
	"embed"

	"github.com/labstack/echo/v4"
)

// SetupRoutes will setup all routes in systems. Sub routs will be manage details by setupXXX function
func SetupRoutes(e *echo.Echo, staticResources embed.FS) {
	setupStaticResourceRoutes(e.Group(""), staticResources)
	g := e.Group("/api")
	setupSecretRoutes(g.Group("/secrets"))
	setupAWSRoutes(g.Group("/aws"))
}
