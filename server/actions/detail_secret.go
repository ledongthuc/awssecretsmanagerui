package actions

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func GetSecretByARN(region, arn string) (secretsmanager.DescribeSecretOutput, error) {
	svc := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(region),
	}))

	input := &secretsmanager.DescribeSecretInput{
		SecretId: aws.String(arn),
	}
	result, err := svc.DescribeSecret(input)
	if err != nil {
		return secretsmanager.DescribeSecretOutput{}, err
	}
	if result == nil {
		return secretsmanager.DescribeSecretOutput{}, errors.New("Can't get secret")
	}
	return *result, nil
}
