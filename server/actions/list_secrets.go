package actions

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var filterNames []string

func GetFilterNames() []string {
	if filterNames != nil {
		return filterNames
	}

	names := os.Getenv("FILTER_NAMES")
	if names == "" {
		filterNames = []string{}
	} else {
		filterNames = strings.Split(names, ",")
	}

	return filterNames
}

func getFilerNamesSecret(listFilterNames []string) *secretsmanager.Filter {
	keyName := "name"
	nameFilters := &secretsmanager.Filter{
		Key:    &keyName,
		Values: make([]*string, len(listFilterNames)),
	}

	for i := range listFilterNames {
		nameFilters.Values[i] = &listFilterNames[i]
	}

	return nameFilters
}

func GetListSecrets(region string) ([]*secretsmanager.SecretListEntry, error) {
	svc := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))
	maxResult := int64(100)

	index := 0
	var token *string
	secrets := []*secretsmanager.SecretListEntry{}
	for ; token != nil || index == 0; index++ {
		result, err := GetAPageSecrets(svc, token, maxResult)
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, result.SecretList...)
		token = result.NextToken
	}
	return secrets, nil
}

func GetAPageSecrets(svc *secretsmanager.SecretsManager, token *string, maxResult int64) (*secretsmanager.ListSecretsOutput, error) {
	input := &secretsmanager.ListSecretsInput{
		MaxResults: &maxResult,
		NextToken:  token,
	}

	filterNames := GetFilterNames()

	var filterNamesSecret *secretsmanager.Filter
	if len(filterNames) > 0 {
		filterNamesSecret = getFilerNamesSecret(filterNames)
	}

	if filterNamesSecret != nil {
		input.Filters = []*secretsmanager.Filter{
			filterNamesSecret,
		}
	}

	result, err := svc.ListSecrets(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}
