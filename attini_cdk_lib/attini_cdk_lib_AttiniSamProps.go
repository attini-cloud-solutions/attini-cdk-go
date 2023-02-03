// Attini CDK Constructs
package attini_cdk_lib


type AttiniSamProps struct {
	Project *Project `field:"required" json:"project" yaml:"project"`
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	Action *string `field:"optional" json:"action" yaml:"action"`
	ConfigFile *string `field:"optional" json:"configFile" yaml:"configFile"`
	EnableTerminationProtection *bool `field:"optional" json:"enableTerminationProtection" yaml:"enableTerminationProtection"`
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	StackRoleArn *string `field:"optional" json:"stackRoleArn" yaml:"stackRoleArn"`
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

