//go:build !no_runtime_type_checking

// Attini CDK Constructs
package attini_cdk_lib

import (
	"fmt"
)

func validateAttiniPayload_StackParameterParameters(parameterName *string) error {
	if parameterName == nil {
		return fmt.Errorf("parameter parameterName is required, but nil was provided")
	}

	return nil
}

