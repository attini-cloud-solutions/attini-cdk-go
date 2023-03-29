// Attini CDK Constructs
package attini_cdk_lib

import (
	_init_ "github.com/attini-cloud-solutions/attini-cdk-go/attini_cdk_lib/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Utility class for reading values from the payload of an Attini deployment plan.
//
// Each method can only be used as the value of a key/value pair in an Attini or
// StepFunction step.
//
// Example:
//   Valid example:
//
//   {my-key: AttiniPayload.environment()}
//
//   Invalid examples:
//
//   {my-key: 'test'+ AttiniPayload.environment()}
//   {my-key: '[step.AttiniPayload.environment()]}
//
type AttiniPayload interface {
}

// The jsii proxy struct for AttiniPayload
type jsiiProxy_AttiniPayload struct {
	_ byte // padding
}

func NewAttiniPayload_Override(a AttiniPayload) {
	_init_.Initialize()

	_jsii_.Create(
		"@attini/cdk.AttiniPayload",
		nil, // no parameters
		a,
	)
}

func AttiniPayload_DistributionName() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniPayload",
		"distributionName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func AttiniPayload_Environment() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniPayload",
		"environment",
		nil, // no parameters
		&returns,
	)

	return returns
}

func AttiniPayload_StackParameter(parameterName *string) *string {
	_init_.Initialize()

	if err := validateAttiniPayload_StackParameterParameters(parameterName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniPayload",
		"stackParameter",
		[]interface{}{parameterName},
		&returns,
	)

	return returns
}

func AttiniPayload_Version() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniPayload",
		"version",
		nil, // no parameters
		&returns,
	)

	return returns
}

func AttiniPayload_DISTRIBUTION_NAME_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniPayload",
		"DISTRIBUTION_NAME_PATH",
		&returns,
	)
	return returns
}

func AttiniPayload_DISTRIBUTION_VERSION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniPayload",
		"DISTRIBUTION_VERSION",
		&returns,
	)
	return returns
}

func AttiniPayload_ENVIRONMENT_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniPayload",
		"ENVIRONMENT_PATH",
		&returns,
	)
	return returns
}

func AttiniPayload_STACK_PARAMETERS_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniPayload",
		"STACK_PARAMETERS_PATH",
		&returns,
	)
	return returns
}

