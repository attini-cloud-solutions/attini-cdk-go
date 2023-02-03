//go:build !no_runtime_type_checking

// Attini CDK Constructs
package attini_cdk_lib

import (
	"fmt"
)

func validatePropsUtil_FixCaseParameters(props interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

