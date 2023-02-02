//go:build no_runtime_type_checking

// Attini resources
package attini_cdk_lib

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttiniRunnerJob) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateAddPrefixParameters(x *string) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (a *jsiiProxy_AttiniRunnerJob) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateAttiniRunnerJob_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateAttiniRunnerJob_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniRunnerJob_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniRunnerJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAttiniRunnerJob_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_AttiniRunnerJob) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewAttiniRunnerJobParameters(scope constructs.Construct, id *string, props *AttiniRunnerJobProps) error {
	return nil
}

