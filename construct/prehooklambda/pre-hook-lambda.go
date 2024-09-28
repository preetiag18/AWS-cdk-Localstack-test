package prehooklambda

import (
	"cdk-localstack/construct/lambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Props struct {
	FnEntry *string
	FnName  string
}

func PreHookLambda(scope constructs.Construct, id string, props *Props) awslambda.Function {
	this := constructs.NewConstruct(scope, &id)

	fn := lambda.Lambda(this, id, &lambda.Props{
		FnEntry: props.FnEntry,
		FnName:  props.FnName,
	})
	addPermissions(fn)
	return fn
}

func addPermissions(fun awslambda.Function) awscdklambdagoalpha.GoFunction {
	fun.AddToRolePolicy(
		awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
			Effect: awsiam.Effect_ALLOW,
			Actions: &[]*string{
				jsii.String("codedeploy:PutLifecycleEventHookExecutionStatus"),
			},
			Resources: &[]*string{jsii.String("*")},
		}),
	)
	return fun
}
