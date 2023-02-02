//go:build no_runtime_type_checking

// Attini resources
package attini_cdk_lib

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttiniState) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateAddPrefixParameters(x *string) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (a *jsiiProxy_AttiniState) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateAttiniState_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateAttiniState_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniState_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniState_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAttiniState_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_AttiniState) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewAttiniStateParameters(scope constructs.Construct, id *string) error {
	return nil
}

