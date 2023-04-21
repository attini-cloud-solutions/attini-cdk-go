// Attini CDK Constructs
package attini_cdk_lib


type Ec2Configuration struct {
	// The instance type of the EC2 instance that will host the Runner.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The AMI to use.
	//
	// Can be specified as an imageId, starting with "ami-", or a short hand name like, AmazonLinux2, AmazonLinux2_arm64 etc.
	// For a complete list, please see the documentation.
	//
	// Will default to AmazonLinux2.
	Ami *string `field:"optional" json:"ami" yaml:"ami"`
	// The instance profile name for the EC2 instance.
	//
	// If omitted, then Attini will create
	// an instance profile with the Runners default role.
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
}

