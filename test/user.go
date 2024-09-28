package test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/stretchr/testify/assert"
)

func assertUser(t *testing.T, res string) {
	assert.Equal(t, "ok", res, "Res is ok")
}

func testUser(ctx context.Context, t *testing.T, awsConfig LocalStackAwsConfig) {
	lambdaClient := lambda.NewFromConfig(awsConfig.Config, func(o *lambda.Options) {
		o.BaseEndpoint = awsConfig.Endpoint
	})
	functionName := "get-users"
	payload := []byte(`{"name": "testuser"}`)
	params := &lambda.InvokeInput{
		FunctionName:   aws.String(functionName),
		Payload:        payload,
		InvocationType: "RequestResponse", // Use "Event" for asynchronous invocation
	}
	resp, err := lambdaClient.Invoke(ctx, params)
	failOnError(t, err, "Lambda function invocation")
	if resp.StatusCode != http.StatusOK {
		failOnError(t, errors.New(string(resp.Payload)), "Lambda response status")
	}
	userData := string(resp.Payload)
	failOnError(t, err, "Lambda response decode")
	assertUser(t, userData)
}
