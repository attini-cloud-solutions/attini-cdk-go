// Attini CDK Constructs
package attini_cdk_lib

import (
	_init_ "github.com/attini-cloud-solutions/attini-cdk-go/attini_cdk_lib/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// The manual approval step will pause the deployment plan and wait for confirmation before continuing.
//
// The deployment can be resumed by running the "attini deploy continue" command with the Attini CLI.
type AttiniManualApproval interface {
	AttiniTask
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	// Continuable states of this Chainable.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	// The tree node.
	Node() constructs.Node
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() awsstepfunctions.State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	Type() *string
	SetType(val *string)
	// Add a paralle branch to this state.
	AddBranch(branch awsstepfunctions.StateGraph)
	// Add a choice branch to this state.
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	// Add a map iterator to this state.
	AddIterator(iteration awsstepfunctions.StateGraph)
	// Add a prefix to the stateId of this state.
	AddPrefix(x *string)
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph awsstepfunctions.StateGraph)
	// Get a string value from the steps output.
	//
	// This method can only be used as the value for a key/value pair in another Attini or
	// StepFunction step.
	//
	// Example:
	//   Valid examples:
	//
	//   {my-key: step.getOutput()}
	//   {my-key: step.getOutput('path','to','key')}
	//
	//   Invalid examples:
	//
	//   {my-key: 'test'+ step.getOutput()}
	//   {my-key: '[step.getOutput()]}
	//
	GetOutput(paths ...*string) *string
	// Get the json path to this steps output.
	//
	// Convenience
	// method that will return a string with the following format
	// $.output.<id>.<path>.<path>
	GetOutputPath(paths ...*string) *string
	// Make the indicated state the default choice transition of this state.
	MakeDefault(def awsstepfunctions.State)
	// Make the indicated state the default transition of this state.
	MakeNext(next awsstepfunctions.State)
	// Go to the indicated state after this state.
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	// Render parallel branches in ASL JSON format.
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	RenderChoices() interface{}
	// Render InputPath/Parameters/OutputPath in ASL JSON format.
	RenderInputOutput() interface{}
	// Render map iterator in ASL JSON format.
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	RenderNextEnd() interface{}
	RenderProps() *map[string]interface{}
	// Render ResultSelector in ASL JSON format.
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	RenderRetryCatch() interface{}
	// Render the state as JSON.
	ToStateJson() *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Allows the state to validate itself.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for AttiniManualApproval
type jsiiProxy_AttiniManualApproval struct {
	jsiiProxy_AttiniTask
}

func (j *jsiiProxy_AttiniManualApproval) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniManualApproval) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewAttiniManualApproval(scope constructs.Construct, id *string) AttiniManualApproval {
	_init_.Initialize()

	if err := validateNewAttiniManualApprovalParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttiniManualApproval{}

	_jsii_.Create(
		"@attini/cdk.AttiniManualApproval",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

func NewAttiniManualApproval_Override(a AttiniManualApproval, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@attini/cdk.AttiniManualApproval",
		[]interface{}{scope, id},
		a,
	)
}

func (j *jsiiProxy_AttiniManualApproval)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AttiniManualApproval)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_AttiniManualApproval)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func AttiniManualApproval_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateAttiniManualApproval_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniManualApproval",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func AttiniManualApproval_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateAttiniManualApproval_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniManualApproval",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func AttiniManualApproval_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateAttiniManualApproval_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniManualApproval",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AttiniManualApproval_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAttiniManualApproval_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniManualApproval",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func AttiniManualApproval_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateAttiniManualApproval_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@attini/cdk.AttiniManualApproval",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (a *jsiiProxy_AttiniManualApproval) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := a.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

func (a *jsiiProxy_AttiniManualApproval) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := a.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (a *jsiiProxy_AttiniManualApproval) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := a.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (a *jsiiProxy_AttiniManualApproval) AddPrefix(x *string) {
	if err := a.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

func (a *jsiiProxy_AttiniManualApproval) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := a.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (a *jsiiProxy_AttiniManualApproval) GetOutput(paths ...*string) *string {
	args := []interface{}{}
	for _, a := range paths {
		args = append(args, a)
	}

	var returns *string

	_jsii_.Invoke(
		a,
		"getOutput",
		args,
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) GetOutputPath(paths ...*string) *string {
	args := []interface{}{}
	for _, a := range paths {
		args = append(args, a)
	}

	var returns *string

	_jsii_.Invoke(
		a,
		"getOutputPath",
		args,
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) MakeDefault(def awsstepfunctions.State) {
	if err := a.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

func (a *jsiiProxy_AttiniManualApproval) MakeNext(next awsstepfunctions.State) {
	if err := a.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

func (a *jsiiProxy_AttiniManualApproval) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := a.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderProps() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"renderProps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniManualApproval) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := a.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

