// Attini CDK Constructs
package attini_cdk_lib

import (
	_init_ "github.com/attini-cloud-solutions/attini-cdk-go/attini_cdk_lib/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Environment variables that are available when deploying a CDK app with Attini.
type AttiniRuntimeVariables interface {
}

// The jsii proxy struct for AttiniRuntimeVariables
type jsiiProxy_AttiniRuntimeVariables struct {
	_ byte // padding
}

func NewAttiniRuntimeVariables_Override(a AttiniRuntimeVariables) {
	_init_.Initialize()

	_jsii_.Create(
		"@attini/cdk.AttiniRuntimeVariables",
		nil, // no parameters
		a,
	)
}

func AttiniRuntimeVariables_DISTRIBUTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniRuntimeVariables",
		"DISTRIBUTION_ID",
		&returns,
	)
	return returns
}

func AttiniRuntimeVariables_DISTRIBUTION_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniRuntimeVariables",
		"DISTRIBUTION_NAME",
		&returns,
	)
	return returns
}

func AttiniRuntimeVariables_ENVIRONMENT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniRuntimeVariables",
		"ENVIRONMENT",
		&returns,
	)
	return returns
}

func AttiniRuntimeVariables_INPUT_FILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniRuntimeVariables",
		"INPUT_FILE_PATH",
		&returns,
	)
	return returns
}

func AttiniRuntimeVariables_OUTPUT_FILE_PATH() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniRuntimeVariables",
		"OUTPUT_FILE_PATH",
		&returns,
	)
	return returns
}

func AttiniRuntimeVariables_STEP_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@attini/cdk.AttiniRuntimeVariables",
		"STEP_NAME",
		&returns,
	)
	return returns
}

