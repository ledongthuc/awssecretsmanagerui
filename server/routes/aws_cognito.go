package routes

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/ledongthuc/awssecretsmanagerui/server/actions"
	"github.com/ledongthuc/awssecretsmanagerui/server/auth"
)

func SetupAWSCognitoRoute(g *echo.Group) {
	g.GET("/auth", func(c echo.Context) error {
		authCode := c.QueryParam("code")
		if len(authCode) == 0 {
			return c.JSON(http.StatusBadRequest, "missing auth code")
		}
		token, err := actions.GetAWSCognitoAccessToken(authCode)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if token.AccessToken == nil {
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("Missing access token"))
		}
		userInfo, err := actions.GetAWSCognitoUserInfo(*token.AccessToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if valid, err := userInfo.IsValidEmail(); !valid {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		h := sha256.New()
		h.Write([]byte(*token.AccessToken))
		clientToken := base64.StdEncoding.EncodeToString(h.Sum(nil))

		cookie := new(http.Cookie)
		cookie.Name = "aws_cognito_token"
		cookie.Value = clientToken
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		auth.GetAWSCognitoLoginUsers().Store(clientToken, struct{}{})
		return c.HTML(http.StatusOK, `<html>
	<head><script>location.href = "/";</script></head>
	<body></body>
</html>`)
		// return c.Redirect(http.StatusTemporaryRedirect, "/")
	})
}
