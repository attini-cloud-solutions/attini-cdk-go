// Attini CDK Constructs
package attini_cdk_lib

import (
	_init_ "github.com/attini-cloud-solutions/attini-cdk-go/attini_cdk_lib/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AttiniRunnerJob is a quick and cost-efficient way to run shell commands using a container.
type AttiniRunnerJob interface {
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
	// Get the json path to this steps output.
	//
	// Convenience
	// method that will return a string with the following format
	// $.output.<id>.<path>.
	GetOutputPath(path *string) *string
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

// The jsii proxy struct for AttiniRunnerJob
type jsiiProxy_AttiniRunnerJob struct {
	jsiiProxy_AttiniTask
}

func (j *jsiiProxy_AttiniRunnerJob) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AttiniRunnerJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewAttiniRunnerJob(scope constructs.Construct, id *string, props *AttiniRunnerJobProps) AttiniRunnerJob {
	_init_.Initialize()

	if err := validateNewAttiniRunnerJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttiniRunnerJob{}

	_jsii_.Create(
		"@attini/cdk.AttiniRunnerJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAttiniRunnerJob_Override(a AttiniRunnerJob, scope constructs.Construct, id *string, props *AttiniRunnerJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@attini/cdk.AttiniRunnerJob",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AttiniRunnerJob)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AttiniRunnerJob)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_AttiniRunnerJob)SetType(val *string) {
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
func AttiniRunnerJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateAttiniRunnerJob_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniRunnerJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func AttiniRunnerJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateAttiniRunnerJob_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniRunnerJob",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func AttiniRunnerJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateAttiniRunnerJob_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniRunnerJob",
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
func AttiniRunnerJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAttiniRunnerJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@attini/cdk.AttiniRunnerJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func AttiniRunnerJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateAttiniRunnerJob_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"@attini/cdk.AttiniRunnerJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := a.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := a.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := a.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) AddPrefix(x *string) {
	if err := a.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := a.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) GetOutputPath(path *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getOutputPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) MakeDefault(def awsstepfunctions.State) {
	if err := a.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) MakeNext(next awsstepfunctions.State) {
	if err := a.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

func (a *jsiiProxy_AttiniRunnerJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
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

func (a *jsiiProxy_AttiniRunnerJob) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) RenderProps() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"renderProps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttiniRunnerJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := a.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

