package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ledongthuc/awssecretsmanagerui/server/actions"
)

func setupAWSRoutes(g *echo.Group) {
	g.GET("/regions", func(c echo.Context) error {
		regions, err := actions.GetAWSRegions()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, regions)
	})

}
