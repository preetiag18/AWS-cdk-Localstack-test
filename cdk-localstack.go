package main

import (
	"cdk-localstack/common/cdk"
	"cdk-localstack/stack"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	dataStack := stack.DataStack(app, "DataStack", &stack.DataStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
	})

	userStack := stack.UserStack(app, "UserStack", &stack.UserStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
	})

	cdk.TagWith("applicationGroup", "", dataStack, userStack)
	cdk.TagWith("application", "web-app", dataStack, userStack)

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("000000000000"),
		Region:  jsii.String("us-east-1"),
	}
}
