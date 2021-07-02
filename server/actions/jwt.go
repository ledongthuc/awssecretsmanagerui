package actions

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginClaims struct {
	jwt.StandardClaims
}

type TokenPayload struct {
	Username string `json:"username"`
}

type Token struct {
	Token   string    `json:"token"`
	Expiry  int       `json:"expiry"`
	Created time.Time `json:"created"`
}

func GenerateJWTToken(data TokenPayload, expiry int) (*Token, error) {
	standardClaims := jwt.StandardClaims{
		IssuedAt: time.Now().UTC().Unix(),
		Subject:  data.Username,
	}

	if expiry != -1 {
		standardClaims.ExpiresAt = time.Now().UTC().Add(time.Second * time.Duration(expiry)).Unix()
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		standardClaims,
	})

	myToken, err := t.SignedString([]byte(os.Getenv("JWT_SIGNATURE_KEY")))
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}
