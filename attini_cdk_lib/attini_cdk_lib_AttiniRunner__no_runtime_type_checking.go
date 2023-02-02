//go:build no_runtime_type_checking

// Attini resources
package attini_cdk_lib

// Building without runtime type checking enabled, so all the below just return nil

func validateAttiniRunner_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAttiniRunnerParameters(scope constructs.Construct, id *string, props *AttiniRunnerProps) error {
	return nil
}

