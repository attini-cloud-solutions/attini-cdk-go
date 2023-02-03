//go:build no_runtime_type_checking

// Attini CDK Constructs
package attini_cdk_lib

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttiniImport) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateAddPrefixParameters(x *string) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (a *jsiiProxy_AttiniImport) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateAttiniImport_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateAttiniImport_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniImport_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniImport_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAttiniImport_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_AttiniImport) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewAttiniImportParameters(scope constructs.Construct, id *string, props *AttiniImportProps) error {
	return nil
}

