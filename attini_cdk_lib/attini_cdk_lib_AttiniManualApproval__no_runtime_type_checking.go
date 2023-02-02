//go:build no_runtime_type_checking

// Attini resources
package attini_cdk_lib

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttiniManualApproval) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateAddPrefixParameters(x *string) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (a *jsiiProxy_AttiniManualApproval) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateAttiniManualApproval_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateAttiniManualApproval_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniManualApproval_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniManualApproval_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAttiniManualApproval_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_AttiniManualApproval) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewAttiniManualApprovalParameters(scope constructs.Construct, id *string) error {
	return nil
}

