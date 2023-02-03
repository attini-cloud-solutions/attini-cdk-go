// Attini CDK Constructs
package attini_cdk_lib


type AttiniCfnProps struct {
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	Template *string `field:"required" json:"template" yaml:"template"`
	Action *string `field:"optional" json:"action" yaml:"action"`
	ConfigFile *string `field:"optional" json:"configFile" yaml:"configFile"`
	EnableTerminationProtection *bool `field:"optional" json:"enableTerminationProtection" yaml:"enableTerminationProtection"`
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	Region *string `field:"optional" json:"region" yaml:"region"`
	StackRoleArn *string `field:"optional" json:"stackRoleArn" yaml:"stackRoleArn"`
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

