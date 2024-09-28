package cdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func TagWith(tagKey string, tagValue string, stacks ...awscdk.Stack) {
	for _, stack := range stacks {
		awscdk.Tags_Of(stack).Add(jsii.String(tagKey), jsii.String(tagValue), &awscdk.TagProps{})
	}
}
