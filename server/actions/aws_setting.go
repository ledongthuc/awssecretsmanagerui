package actions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/endpoints"
)

type AWSSetting struct {
	Region string
}

func GetAWSSetting() AWSSetting {
	return AWSSetting{
		Region: "eu-north-1",
	}
}

func GetAWSRegions() ([]string, error) {
	resolver := endpoints.DefaultResolver()
	epartitions, ok := resolver.(endpoints.EnumPartitions)
	if !ok {
		return []string{}, fmt.Errorf("can't load and resolve regions")
	}

	partitions := epartitions.Partitions()
	result := []string{}
	for _, p := range partitions {
		for region := range p.Regions() {
			result = append(result, region)
		}
	}
	return result, nil
}
