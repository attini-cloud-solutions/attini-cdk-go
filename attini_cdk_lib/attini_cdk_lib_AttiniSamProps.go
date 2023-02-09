// Attini CDK Constructs
package attini_cdk_lib


type AttiniSamProps struct {
	// SAM project config.
	Project *Project `field:"required" json:"project" yaml:"project"`
	// The name that should be given to the stack when deployed.
	//
	// The name must be unique in the Region in which you are creating the stack.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Specify if the stack should be created/updated or deleted.
	Action CfnAction `field:"optional" json:"action" yaml:"action"`
	// Specifies a path to a configuration file for the stack.
	//
	// For more information, see the {@link https://docs.attini.io/api-reference/cloudformation-configuration.html#api-reference-cloudformation-configuration documentation}
	ConfigFile *string `field:"optional" json:"configFile" yaml:"configFile"`
	// Specify if termination protection should be enabled for the stack.
	//
	// For more information, see the {@link https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-protect-stacks.html AWS documentation}.
	EnableTerminationProtection *bool `field:"optional" json:"enableTerminationProtection" yaml:"enableTerminationProtection"`
	// The role that should be assumed when the Attini Framework deploys the CloudFormation stack.
	//
	// The ExecutionRole has to trust the following role so that it can be assumed:
	// ```
	// arn:aws:iam::{AccountId}:role/attini/attini-action-role-{Region}
	// ```.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The parameters for the stack.
	//
	// For more information, see the {@link https://docs.attini.io/api-reference/cloudformation-configuration.html#api-reference-cloudformation-configuration documentation}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The arn of the StackRole, find more info here: {@link https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-servicerole.html AWS CloudFormation service role}.
	StackRoleArn *string `field:"optional" json:"stackRoleArn" yaml:"stackRoleArn"`
	// The tags for the stack.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Variables that should be passed to any {@link configFile} that is configured.
	//
	// Variables can be referenced in the configuration file and can be used to pass data from the payload to the configuration.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

