package bauth

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	BcryptCost int
	Users      Users `json:"users"`
}

func LoadConfig(url string) (AuthConfig, error) {
	f, err := os.Open(url)
	if err != nil {
		return AuthConfig{}, errors.Wrap(err, "load file")
	}
	dec := json.NewDecoder(f)
	var c AuthConfig
	if err := dec.Decode(&c); err != nil {
		return AuthConfig{}, errors.Wrap(err, "parse file")
	}

	c.FillDefault()
	return c, nil
}

func (c *AuthConfig) FillDefault() {
	if c == nil {
		return
	}
	if c.BcryptCost == 0 {
		c.BcryptCost = bcrypt.DefaultCost
	}
}
