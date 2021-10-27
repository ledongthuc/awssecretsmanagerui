package actions

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetAWSCognitoAppName() string {
	return os.Getenv("AWS_COGNITO_APP_NAME")
}

func GetAWSCognitoRegion() string {
	return os.Getenv("AWS_COGNITO_REGION")
}

func GetAWSCognitoClientID() string {
	return os.Getenv("AWS_COGNITO_CLIENT_ID")
}

func GetAWSCognitoClientSecret() string {
	return os.Getenv("AWS_COGNITO_CLIENT_SECRET")
}

func GetAWSCognitoHost() string {
	return fmt.Sprintf(
		"https://%s.auth.%s.amazoncognito.com",
		GetAWSCognitoAppName(),
		GetAWSCognitoRegion(),
	)
}

type CognitoToken struct {
	AccessToken *string `json:"access_token"`
	TokenType   *string `json:"token_type"`
	IDToken     *string `json:"id_token"`
	ExpiresIn   *int64  `json:"expires_in"`
}

func GetAWSCognitoAccessToken(authCode string) (CognitoToken, error) {
	urlStr := GetAWSCognitoHost() + "/oauth2/token"
	request := url.Values{}
	request.Set("grant_type", "authorization_code")
	request.Set("client_id", GetAWSCognitoClientID())
	request.Set("code", authCode)
	request.Set("redirect_uri", os.Getenv("AWS_COGNITO_REDIRECT_URL"))
	r, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(request.Encode()))
	if err != nil {
		return CognitoToken{}, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", GetAWSCognitoClientID(), GetAWSCognitoClientSecret())))))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return CognitoToken{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CognitoToken{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return CognitoToken{}, errors.New(string(body))
	}

	var token CognitoToken
	if err := json.Unmarshal(body, &token); err != nil {
		return CognitoToken{}, err
	}

	return token, nil
}

type CognitoUserInfo struct {
	Sub               string `json:"sub"`
	Name              string `json:"name"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	PreferredUsername string `json:"preferred_username"`
	Email             string `json:"email"`
}

func (c CognitoUserInfo) IsValidEmail() (bool, error) {
	emails := os.Getenv("AWS_COGNITO_ALLOWED_EMAILS")
	if len(emails) == 0 {
		return true, nil
	}
	for _, email := range strings.Split(emails, ",") {
		if email == c.Email {
			return true, nil
		}
	}
	return false, fmt.Errorf("%s doesn't have permission to access system", c.Email)
}

func GetAWSCognitoUserInfo(accessToken string) (CognitoUserInfo, error) {
	urlStr := GetAWSCognitoHost() + "/oauth2/userInfo"
	r, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return CognitoUserInfo{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CognitoUserInfo{}, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return CognitoUserInfo{}, errors.New(string(body))
	}

	var userInfo CognitoUserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return CognitoUserInfo{}, err
	}
	return userInfo, nil
}
