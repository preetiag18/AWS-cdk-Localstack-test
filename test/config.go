package test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

type LocalStackAwsConfig struct {
    Config aws.Config
    Endpoint *string
}

func getAwsConfig(ctx context.Context, t *testing.T, localstack *testcontainers.DockerContainer) LocalStackAwsConfig{
    endpoint, err := localstack.PortEndpoint(ctx, nat.Port("4566"), "http")
    failOnError(t, err, "Localstack endpoint")
    awsConfig, err := config.LoadDefaultConfig(ctx,
        config.WithRegion("us-east-1"),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "test")),
    )
    failOnError(t, err, "Load AWS config")
    return LocalStackAwsConfig{
        Config: awsConfig,
        Endpoint: aws.String(endpoint),
    }
}