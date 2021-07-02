package auth

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ledongthuc/awssecretsmanagerui/server/actions"
)

func CreateJWTAuth() middleware.JWTConfig {
	return middleware.JWTConfig{
		TokenLookup:  "headers:authorization",
		SigningKey:   []byte(os.Getenv("JWT_SIGNATURE_KEY")),
		Claims:       &actions.LoginClaims{},
		ErrorHandler: JWTErrorChecker,
	}
}

func JWTErrorChecker(err error) error {
	echoError, ok := err.(*echo.HTTPError)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return echo.NewHTTPError(http.StatusUnauthorized, echoError.Message)
}
