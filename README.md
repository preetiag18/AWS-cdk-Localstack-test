# AWS Infra (CDK) testing using Localstack

This demo project explains how to use LocalStack for testing AWS infrastructure locally and in CI pipelines.

## Infrastructure as a Code (IaaC)

The AWS infrastructure for this project is written using the [AWS CDK in Go](https://docs.aws.amazon.com/cdk/v2/guide/work-with-cdk-go.html)

### Stacks

This project consists of two AWS CDK stacks:

1. Lambda Stack: A Lambda function that returns a simple string response.
2. SQS Stack: An SQS queue to send and receive data.

### Testing

#### Infrastrure testing

1. The project uses LocalStack (running as a Docker container) to simulate AWS services locally.
2. The AWS CDK infrastructure is then deployed using `cdklocal deploy`.
3. After deployment, the Lambda function is invoked using the AWS SDK.
4. Assertions are performed on the Lambda response to validate the functionality.

For more details, refer to the test file: [test/cdk-localstack_e2e_test.go](./test/cdk-localstack_e2e_test.go)

In this way we can test both the infrastructure code and the Lambda business logic.

You can run the end-to-end (E2E) infrastructure tests with the following command:

```bash
go test -timeout 30s -run ^TestApp$ cdk-localstack/test
```

#### Template assertions

We can also write unit tests to validate that the generated CloudFormation template contains the expected resources with the correct properties.

For more details, refer to the test file: [cdk-localstack_test.go](./cdk-localstack_test.go)

To run the unit tests for template assertions, use the following command:

```bash
go test -timeout 30s -run ^TestDataStack$ cdk-localstack
```
