package routes

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ledongthuc/awssecretsmanagerui/server/actions"
)

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func SetupLoginRoute(g *echo.Group) {
	g.POST("/login", login)
}

func login(c echo.Context) error {
	var loginData UserLogin

	if err := c.Bind(&loginData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if ok, err := actions.ValidPassword(loginData.Username, loginData.Password); !ok {
		if err == nil {
			err = errors.New("username or password is incorrect")
		}
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	payload := actions.TokenPayload{
		Sub:      loginData.Username,
		Username: loginData.Username,
	}

	expiry, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_TIME"))
	if err != nil {
		expiry = -1
	}

	token, err := actions.GenerateJWTToken(payload, expiry)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
