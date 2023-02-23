// Attini CDK Constructs
package attini_cdk_lib


type StackConfigurationProps struct {
	// Parameter configuration for the CloudFormation stack/stacks.
	Parameters *map[string]*string `field:"required" json:"parameters" yaml:"parameters"`
	// Stack name that you want to configure.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
}

