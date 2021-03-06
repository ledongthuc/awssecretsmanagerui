package actions

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/endpoints"
)

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
