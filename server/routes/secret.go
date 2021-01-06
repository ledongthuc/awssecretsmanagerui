package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ledongthuc/awssecretsmanagerui/server/actions"
)

func setupSecretRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		secrets, err := actions.GetListSecrets()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secrets)
	})
}
