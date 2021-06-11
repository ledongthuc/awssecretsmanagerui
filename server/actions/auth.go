package actions

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Accounts map[string]string

func GetUsers() (Accounts, error) {
	raw := os.Getenv("AUTH_ACCOUNTS")
	var accounts Accounts
	err := json.Unmarshal([]byte(raw), &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func ValidPassword(username string, password string) (bool, error) {
	accounts, err := GetUsers()
	if err != nil {
		return false, errors.Wrap(err, "can't load accounts")
	}
	hashedPassword, existed := accounts[username]
	if !existed {
		return false, nil
	}

	if fmt.Sprintf("%x", sha256.Sum256([]byte(password))) != hashedPassword {
		return false, nil
	}
	return true, nil
}
