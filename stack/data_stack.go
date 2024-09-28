package stack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DataStackProps struct {
	awscdk.StackProps
}

func DataStack(scope constructs.Construct, id string, props *DataStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awssqs.NewQueue(stack, jsii.String("CdkLocalstackQueue"), &awssqs.QueueProps{
		QueueName:         jsii.String("local-queue"),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	})

	return stack
}
