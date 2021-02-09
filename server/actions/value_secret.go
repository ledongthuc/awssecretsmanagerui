package actions

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetSecretValueByARN(region, arn string) (secretsmanager.GetSecretValueOutput, error) {
	svc := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(arn),
	}
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return secretsmanager.GetSecretValueOutput{}, err
	}
	if result == nil {
		return secretsmanager.GetSecretValueOutput{}, errors.New("Can't get secret")
	}
	return *result, nil
}

func UpdateSecretValue(region string, request secretsmanager.PutSecretValueInput) (secretsmanager.GetSecretValueOutput, error) {
	svc := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	_, err := svc.PutSecretValue(&request)
	if err != nil {
		return secretsmanager.GetSecretValueOutput{}, err
	}
	return GetSecretValueByARN(region, *request.SecretId)
}

func UpdateSecretValueBinary(region string, arn string, binaryVaue []byte) (secretsmanager.GetSecretValueOutput, error) {
	svc := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	request := secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(arn),
		SecretBinary: binaryVaue,
	}
	_, err := svc.PutSecretValue(&request)
	if err != nil {
		return secretsmanager.GetSecretValueOutput{}, err
	}
	return GetSecretValueByARN(region, arn)
}
