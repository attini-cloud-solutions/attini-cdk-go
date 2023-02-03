//go:build !no_runtime_type_checking

// Attini resources
package attini_cdk_lib

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DeploymentPlanStack) validateAddDependencyParameters(target awscdk.Stack) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateAddTransformParameters(transform *string) error {
	if transform == nil {
		return fmt.Errorf("parameter transform is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateAllocateLogicalIdParameters(cfnElement awscdk.CfnElement) error {
	if cfnElement == nil {
		return fmt.Errorf("parameter cfnElement is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateExportValueParameters(exportedValue interface{}, options *awscdk.ExportValueOptions) error {
	if exportedValue == nil {
		return fmt.Errorf("parameter exportedValue is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateFormatArnParameters(components *awscdk.ArnComponents) error {
	if components == nil {
		return fmt.Errorf("parameter components is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(components, func() string { return "parameter components" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateGetLogicalIdParameters(element awscdk.CfnElement) error {
	if element == nil {
		return fmt.Errorf("parameter element is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateRenameLogicalIdParameters(oldId *string, newId *string) error {
	if oldId == nil {
		return fmt.Errorf("parameter oldId is required, but nil was provided")
	}

	if newId == nil {
		return fmt.Errorf("parameter newId is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateReportMissingContextKeyParameters(report *cloudassemblyschema.MissingContext) error {
	if report == nil {
		return fmt.Errorf("parameter report is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(report, func() string { return "parameter report" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateResolveParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateSplitArnParameters(arn *string, arnFormat awscdk.ArnFormat) error {
	if arn == nil {
		return fmt.Errorf("parameter arn is required, but nil was provided")
	}

	if arnFormat == "" {
		return fmt.Errorf("parameter arnFormat is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeploymentPlanStack) validateToJsonStringParameters(obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("parameter obj is required, but nil was provided")
	}

	return nil
}

func validateDeploymentPlanStack_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateDeploymentPlanStack_IsStackParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateDeploymentPlanStack_OfParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateNewDeploymentPlanStackParameters(scope constructs.Construct, id *string, props *awscdk.StackProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

