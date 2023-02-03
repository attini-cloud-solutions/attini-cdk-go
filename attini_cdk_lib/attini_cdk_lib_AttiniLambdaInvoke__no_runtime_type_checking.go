//go:build no_runtime_type_checking

// Attini CDK Constructs
package attini_cdk_lib

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AttiniLambdaInvoke) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateAddPrefixParameters(x *string) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (a *jsiiProxy_AttiniLambdaInvoke) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateAttiniLambdaInvoke_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateAttiniLambdaInvoke_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniLambdaInvoke_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateAttiniLambdaInvoke_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAttiniLambdaInvoke_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_AttiniLambdaInvoke) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewAttiniLambdaInvokeParameters(scope constructs.Construct, id *string, props *AttiniLambdaInvokeProps) error {
	return nil
}

