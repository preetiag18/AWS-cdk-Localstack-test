package test

import (
	"context"
	"testing"
)

func TestApp(t *testing.T) {
	ctx := context.Background()
	services := startServices(ctx, t) // starts localstack docker container
	deploy(t)                                       // deploys infrastructure on localstack
	awsConfig := getAwsConfig(ctx, t, services.localstack)
	testSqsEventSending(ctx, t, awsConfig) // invoke test assertions
	testUser(ctx, t, awsConfig)            // invoke test assertions
}
