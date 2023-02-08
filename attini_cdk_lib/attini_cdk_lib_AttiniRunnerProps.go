// Attini CDK Constructs
package attini_cdk_lib


type AttiniRunnerProps struct {
	// VPC configuration.
	AwsVpcConfiguration *AwsVpcConfiguration `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
	// The name of the container in the task definition.
	//
	// This is only required if a task definition with multiple containers is specified.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The name of the ECS Cluster to use.
	//
	// This is not mandatory if there is a default cluster in the account.
	EcsCluster *string `field:"optional" json:"ecsCluster" yaml:"ecsCluster"`
	// A Container image that the runner should use.
	//
	// If you configure this value, Attini will configure a TaskDefinition for you.
	//
	// This configuration can not be combined with the {@link taskDefinitionArn} configuration.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The IAM Role the Runner should use.
	//
	// This IAM Role will override the IAM Role from the TaskDefinition.
	//
	// This IAM Role requires a basic execution policy that allows the runner to communicate with the deployment plan.
	// Add the following execution policy to the IAM Role:
	// ```
	// arn:aws:iam::${AccountId}:policy/attini-runner-basic-execution-policy-${Region}
	// ```.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Configuration for the runner.
	RunnerConfiguration *RunnerConfiguration `field:"optional" json:"runnerConfiguration" yaml:"runnerConfiguration"`
	// Runner startup configuration.
	Startup *Startup `field:"optional" json:"startup" yaml:"startup"`
	// Fargate ECS task definition that the Attini Runner should use.
	//
	// If omitted then Attini will use its default task definition.
	TaskDefinitionArn *string `field:"optional" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
}

