// Attini resources
package attini_cdk_lib


type AttiniRunnerProps struct {
	AwsVpcConfiguration *AwsVpcConfiguration `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	EcsCluster *string `field:"optional" json:"ecsCluster" yaml:"ecsCluster"`
	Image *string `field:"optional" json:"image" yaml:"image"`
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	RunnerConfiguration *RunnerConfiguration `field:"optional" json:"runnerConfiguration" yaml:"runnerConfiguration"`
	Startup *Startup `field:"optional" json:"startup" yaml:"startup"`
	TaskDefinitionArn *string `field:"optional" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
}

