package auth

import (
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var awsCognitoLoginUsers *sync.Map

func GetAWSCognitoLoginUsers() *sync.Map {
	if awsCognitoLoginUsers == nil {
		awsCognitoLoginUsers = &sync.Map{}
	}
	return awsCognitoLoginUsers
}

func AWSCognitoMiddleware(skipPathes []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, path := range skipPathes {
				if strings.HasPrefix(c.Request().URL.Path, path) {
					return next(c)
				}
			}
			cookie, err := c.Cookie("aws_cognito_token")
			if err != nil {
				return awsCognitoError(c, err.Error())
			}
			if cookie == nil {
				return awsCognitoError(c, "Fail to load cookie")
			}

			_, valid := GetAWSCognitoLoginUsers().Load(cookie.Value)
			if !valid {
				return awsCognitoError(c, "invalid or expired token")
			}

			return next(c)
		}
	}
}

func awsCognitoError(c echo.Context, errMsg string) error {
	loginURL := os.Getenv("AWS_COGNITO_LOGIN_URL")
	if len(loginURL) == 0 {
		return echo.NewHTTPError(http.StatusUnauthorized, errMsg)
	}
	log.Warn("Unauthentication: ", errMsg)
	return c.Redirect(http.StatusTemporaryRedirect, loginURL)
}
