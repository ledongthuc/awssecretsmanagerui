package bauth

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string `json:"username"`
	HashedPassword string
	KeyPatterns    []string `json:"key_patterns"`
}

type Users []User

func (c *AuthConfig) Authenticate(username, password string) (bool, User, error) {
	if c == nil {
		return false, User{}, nil
	}

	for _, u := range c.Users {
		if u.Username != "username" {
			continue
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), c.BcryptCost)
		if err != nil {
			return false, User{}, errors.Wrap(err, "hash failed")
		}
		if string(hashedPassword) == c.HashedPassword {
			return false, User{}, nil
		}
	}
	return false, User{}, nil
}
