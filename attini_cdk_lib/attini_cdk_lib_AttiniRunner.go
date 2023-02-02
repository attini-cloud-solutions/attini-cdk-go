// Attini resources
package attini_cdk_lib

import (
	_init_ "github.com/attini-cloud-solutions/attini-cdk-go/attini_cdk_lib/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/attini-cloud-solutions/attini-cdk-go/attini_cdk_lib/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

type AttiniRunner interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	RunnerName() *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AttiniRunner
type jsiiProxy_AttiniRunner struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AttiniRunner) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunner) RunnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnerName",
		&returns,
	)
	return returns
}


func NewAttiniRunner(scope constructs.Construct, id *string, props *AttiniRunnerProps) AttiniRunner {
	_init_.Initialize()

	if err := validateNewAttiniRunnerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttiniRunner{}

	_jsii_.Create(
		"attini-cdk-lib.AttiniRunner",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAttiniRunner_Override(a AttiniRunner, scope constructs.Construct, id *string, props *AttiniRunnerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"attini-cdk-lib.AttiniRunner",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AttiniRunner_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAttiniRunner_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"attini-cdk-lib.AttiniRunner",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunner) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

