package routes

import (
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/labstack/echo/v4"

	"github.com/ledongthuc/awssecretsmanagerui/server/actions"
)

func setupSecretRoutes(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		region := c.QueryParam("region")
		if len(region) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing region")
		}
		secrets, err := actions.GetListSecrets(region)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secrets)
	})

	g.GET("/detail", func(c echo.Context) error {
		arn := c.QueryParam("arn")
		if len(arn) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing selected ARN")
		}
		region := c.QueryParam("region")
		if len(region) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing region")
		}
		secret, err := actions.GetSecretByARN(region, arn)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secret)
	})

	g.GET("/value", func(c echo.Context) error {
		arn := c.QueryParam("arn")
		if len(arn) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing selected ARN")
		}
		region := c.QueryParam("region")
		if len(region) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing region")
		}
		secret, err := actions.GetSecretValueByARN(region, arn)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secret)
	})

	g.POST("/value/update", func(c echo.Context) error {
		region := c.QueryParam("region")
		if len(region) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing region")
		}

		var request secretsmanager.PutSecretValueInput
		err := c.Bind(&request)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		secret, err := actions.UpdateSecretValue(region, request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secret)
	})

	g.POST("/value/upload", func(c echo.Context) error {
		arn := c.QueryParam("arn")
		if len(arn) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing selected ARN")
		}

		region := c.QueryParam("region")
		if len(region) == 0 {
			return echo.NewHTTPError(http.StatusInternalServerError, "Missing region")
		}

		file, err := c.FormFile("file")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		src, err := file.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()
		b, err := ioutil.ReadAll(src)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		secret, err := actions.UpdateSecretValueBinary(region, arn, b)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, secret)
	})
}
