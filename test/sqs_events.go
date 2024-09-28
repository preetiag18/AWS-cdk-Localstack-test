package test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/jsii-runtime-go"

	_ "github.com/aws/jsii-runtime-go"
)

func GetQueueURL(c context.Context, client *sqs.Client, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return client.GetQueueUrl(c, input)
}

func testSqsEventSending(ctx context.Context, t *testing.T, awsConfig LocalStackAwsConfig) {
	client := sqs.NewFromConfig(awsConfig.Config, func(o *sqs.Options) {
		o.BaseEndpoint = awsConfig.Endpoint
	})
	queueInput := &sqs.GetQueueUrlInput{
		QueueName:              jsii.String("local-queue"),
		QueueOwnerAWSAccountId: jsii.String("000000000000"),
	}
	queue, err := GetQueueURL(ctx, client, queueInput)
	failOnError(t, err, "Get SQS queue URL")
	_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: jsii.String("Dummy message"),
		QueueUrl:    queue.QueueUrl,
	})
	failOnError(t, err, "Event sent to SQS queue")
}
