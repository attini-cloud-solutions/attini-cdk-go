// Attini CDK Constructs
package attini_cdk_lib


// The VPC configuration for the ECS task.
//
// If awsVpcConfiguration is omitted, Attini will use the default VPC and create a new security group resource in the init deploy stack.
// The security group will have no inbound rules (no openings), but allow all outgoing traffic.
type AwsVpcConfiguration struct {
	// Whether the task's elastic network interface receives a public IP address.
	//
	// Default is false.
	// *
	// For more information see the {@link https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_AwsVpcConfiguration.html AWS ECS VPC documentation}
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// A list of the security group ids associated with the ECS task.
	//
	// For more information see the {@link https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_AwsVpcConfiguration.html AWS ECS VPC documentation}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A list of the subnet ids associated with the ECS task.
	//
	// For more information see the AWS ECS VPC documentation.
	//
	// For more information see the {@link https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_AwsVpcConfiguration.html AWS ECS VPC documentation}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

