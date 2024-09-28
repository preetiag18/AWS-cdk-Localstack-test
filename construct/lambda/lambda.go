package lambda

import (
	"cdk-localstack/common/cdk"
	"cdk-localstack/common/constants"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodedeploy"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Props struct {
	EnvironmentVariables *map[string]*string
	FnEntry              *string
	FnName               string
	MemorySize           *float64
	Timeout              awscdk.Duration
	Vpc                  *awsec2.IVpc
	PreHookFn            awslambda.Function
}

func Lambda(scope constructs.Construct, id string, props *Props) awslambda.Function {
	this := constructs.NewConstruct(scope, &id)
	var fn awscdklambdagoalpha.GoFunction

	if props.Vpc != nil {
		fn = awscdklambdagoalpha.NewGoFunction(this, jsii.String(props.FnName), &awscdklambdagoalpha.GoFunctionProps{
			FunctionName: jsii.String(props.FnName),
			Entry:        props.FnEntry,
			Environment:  enrichEnvList(props.EnvironmentVariables),
			MemorySize:   getMemorySize(props.MemorySize),
			Timeout:      getTimeout(props.Timeout),
			Vpc:          *props.Vpc,
		})
	} else {
		fn = awscdklambdagoalpha.NewGoFunction(this, jsii.String(props.FnName), &awscdklambdagoalpha.GoFunctionProps{
			FunctionName: jsii.String(props.FnName),
			Entry:        props.FnEntry,
			Environment:  enrichEnvList(props.EnvironmentVariables),
			MemorySize:   getMemorySize(props.MemorySize),
			Timeout:      getTimeout(props.Timeout),
		})
	}

	if props.PreHookFn != nil {
		version := fn.CurrentVersion()

		props.PreHookFn.AddEnvironment(jsii.String(cdk.LambdaArnToInvoke), version.FunctionArn(), &awslambda.EnvironmentOptions{})

		alias := awslambda.NewAlias(this, jsii.String("LambdaAlias"), &awslambda.AliasProps{
			AliasName: jsii.String("live"),
			Version:   version,
		})

		awscodedeploy.NewLambdaDeploymentGroup(this, jsii.String("DeploymentGroup"), &awscodedeploy.LambdaDeploymentGroupProps{
			PreHook:          props.PreHookFn,
			Alias:            alias,
			DeploymentConfig: awscodedeploy.LambdaDeploymentConfig_ALL_AT_ONCE(),
		})
	}

	return fn
}

func enrichEnvList(baseEnvList *map[string]*string) *map[string]*string {
	if baseEnvList == nil {
		envVariables := make(map[string]*string)
		return &envVariables
	}

	value := constants.EnvPrefixDefaultValue

	(*baseEnvList)[constants.EnvPrefix] = jsii.String(value)
	return baseEnvList
}

func getTimeout(selectedTimeout awscdk.Duration) awscdk.Duration {
	if selectedTimeout == nil {
		return awscdk.Duration_Seconds(jsii.Ptr(float64(60)))
	}
	return selectedTimeout
}

func getMemorySize(selectedMemorySize *float64) *float64 {
	if selectedMemorySize == nil {
		return jsii.Ptr(float64(128))
	}
	return selectedMemorySize
}
