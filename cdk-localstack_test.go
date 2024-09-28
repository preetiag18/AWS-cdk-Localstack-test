package main

import (
	stacks "cdk-localstack/stack"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestDataStack(t *testing.T) {
	app := awscdk.NewApp(nil)

	stack := stacks.DataStack(app, "dataStack", nil)

	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"VisibilityTimeout": 300,
	})
}
