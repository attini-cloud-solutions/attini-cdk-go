// Attini CDK Constructs
package attini_cdk_lib


type AwsVpcConfiguration struct {
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

