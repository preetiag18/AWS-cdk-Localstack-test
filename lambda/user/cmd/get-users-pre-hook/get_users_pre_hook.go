package main

import (
	"cdk-localstack/common/cdk"
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	lambdaSdk "github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/jsii-runtime-go"
	"github.com/rs/zerolog/log"
)

type PreHookApp struct {
	codeDeployClient *codedeploy.Client
	lambdaClient     *lambdaSdk.Client
}

type LambdaDeploymentPreHookInput struct {
	DeploymentID                  string `json:"deploymentId"`
	LifecycleEventHookExecutionID string `json:"lifecycleEventHookExecutionId"`
}

func validateResponse(res string) bool {
	if res == "ok" {
		return true
	}
	return true
}

func (app *PreHookApp) handle(ctx context.Context, event LambdaDeploymentPreHookInput) (*string, error) {
	log.Info().Msg("BeforeAllowTraffic hook tests started")

	resp, err := app.lambdaClient.Invoke(ctx, &lambdaSdk.InvokeInput{
		FunctionName: jsii.String(os.Getenv(cdk.LambdaArnToInvoke)),
		Payload:      []byte(`{"name": "test"}`),
	})

	if err != nil {
		return nil, err
	}

	res := string(resp.Payload)

	preHookValidationStatus := types.LifecycleEventStatusFailed

	if validateResponse(res) {
		preHookValidationStatus = types.LifecycleEventStatusSucceeded
	}

	// Complete the PreTraffic Hook by sending CodeDeploy the validation status
	statusPayload, err := app.codeDeployClient.PutLifecycleEventHookExecutionStatus(ctx,
		&codedeploy.PutLifecycleEventHookExecutionStatusInput{
			DeploymentId:                  jsii.String(event.DeploymentID),
			LifecycleEventHookExecutionId: jsii.String(event.LifecycleEventHookExecutionID),
			Status:                        preHookValidationStatus,
		},
	)
	if err != nil {
		return nil, err
	}

	return statusPayload.LifecycleEventHookExecutionId, nil
}

func main() {
	log.Info().Msg("Initializing lambda deployment pre commit hook fn")

	ctx := context.Background()
	sdkConfig, err := awsConfig.LoadDefaultConfig(ctx)
	if err != nil {
		log.Err(err).Msg("Failed to load default configuration")
		return
	}
	lambdaClient := lambdaSdk.NewFromConfig(sdkConfig)
	codeDeployClient := codedeploy.NewFromConfig(sdkConfig)

	app := PreHookApp{
		codeDeployClient: codeDeployClient,
		lambdaClient:     lambdaClient,
	}

	lambda.Start(app.handle)
}
