package stack

import (
	"cdk-localstack/construct/lambda"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type UserStackProps struct {
    awscdk.StackProps
}

func UserStack(scope constructs.Construct, id string, props *UserStackProps) awscdk.Stack {
    var sprops awscdk.StackProps
    if props != nil {
        sprops = props.StackProps
    }
    stack := awscdk.NewStack(scope, &id, &sprops)

    createGetUsersFn(stack, props)

    return stack
}

func createGetUsersFn(stack awscdk.Stack, props *UserStackProps) {
    lambda.Lambda(stack, "get-users", &lambda.Props{
        FnEntry: jsii.String("lambda/user/cmd/get-users"),
        FnName:  "get-users",
    })
}
