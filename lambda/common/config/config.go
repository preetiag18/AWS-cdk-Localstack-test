package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig(ctx context.Context) (aws.Config, error) {
	if IsLocal() {
		return config.LoadDefaultConfig(ctx,
			config.WithRegion("us-east-1"),
		)
	}
	return config.LoadDefaultConfig(ctx)
}

func IsLocal() bool {
	_, exist := os.LookupEnv("LOCALSTACK_HOSTNAME")
	return exist
}
