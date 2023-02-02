// Attini resources
package attini_cdk_lib

import (
	_init_ "github.com/attini-cloud-solutions/attini-cdk-go/attini_cdk_lib/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type PropsUtil interface {
}

// The jsii proxy struct for PropsUtil
type jsiiProxy_PropsUtil struct {
	_ byte // padding
}

func NewPropsUtil() PropsUtil {
	_init_.Initialize()

	j := jsiiProxy_PropsUtil{}

	_jsii_.Create(
		"attini-cdk-lib.PropsUtil",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewPropsUtil_Override(p PropsUtil) {
	_init_.Initialize()

	_jsii_.Create(
		"attini-cdk-lib.PropsUtil",
		nil, // no parameters
		p,
	)
}

func PropsUtil_FixCase(props interface{}) *map[string]interface{} {
	_init_.Initialize()

	if err := validatePropsUtil_FixCaseParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.StaticInvoke(
		"attini-cdk-lib.PropsUtil",
		"fixCase",
		[]interface{}{props},
		&returns,
	)

	return returns
}

