// Attini CDK Constructs
package attini_cdk_lib


type AttiniImportProps struct {
	SourceType SourceType `field:"required" json:"sourceType" yaml:"sourceType"`
	DistributionSource *DistributionSource `field:"optional" json:"distributionSource" yaml:"distributionSource"`
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	Mapping *map[string]*string `field:"optional" json:"mapping" yaml:"mapping"`
	S3Source *S3Source `field:"optional" json:"s3Source" yaml:"s3Source"`
}

