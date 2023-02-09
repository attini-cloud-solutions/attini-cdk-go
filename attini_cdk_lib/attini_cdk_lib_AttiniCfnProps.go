// Attini CDK Constructs
package attini_cdk_lib


type AttiniCfnProps struct {
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
	// An optional field name that the CloudFormations output should be placed under in the deployment plan payload.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// The CloudFormation parameters.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The region that the template should be deployed to.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The name that should be given to the stack when deployed.
	//
	// The name must be unique in the Region in which you are creating the stack.
	//
	// Required if it is not specified in {@link configFile}.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// The arn of the StackRole, find more info here: {@link https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-servicerole.html AWS CloudFormation service role}.
	StackRoleArn *string `field:"optional" json:"stackRoleArn" yaml:"stackRoleArn"`
	// The CloudFormation tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The path to the CloudFormation template. Can either be:.
	//
	// 1. A path to a file in the distribution. The path should be from the root of the project and start with a "/".
	// 2. A URL to a public S3 file, starting with "https://".
	// 3. An S3 path, starting with "s3://".
	//
	// Required if it is not specified in {@link configFile}.
	Template *string `field:"optional" json:"template" yaml:"template"`
	// Variables that should be passed to any {@link configFile} that is configured.
	//
	// Variables can be referenced in the configuration file and can be used to pass data from the payload to the configuration.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

