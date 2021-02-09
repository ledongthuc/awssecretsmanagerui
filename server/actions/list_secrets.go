package actions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

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
	result, err := svc.ListSecrets(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}
