// Attini CDK Constructs
package attini_cdk_lib


type AttiniImportProps struct {
	// Specifies what kind of source should be used.
	SourceType SourceType `field:"required" json:"sourceType" yaml:"sourceType"`
	// Used when the source of the import should be another distribution deployed in the environment.
	//
	// Before a distribution can import the output of another distribution, it first needs to be declared as a dependency in the attini-configuration file.
	DistributionSource *DistributionSource `field:"optional" json:"distributionSource" yaml:"distributionSource"`
	// The arn of the execution role that should be used for accessing the source.
	//
	// At the moment only needed for the S3 source type if Attini does not have access to the S3 Bucket.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// A key/value map where the value is a path to a value in the imported document.
	//
	// The path follows the {@link https://goessner.net/ JSONPath} syntax. The value on the path will be included in the output of the step under the same key name as the mapping.
	Mapping *map[string]*string `field:"optional" json:"mapping" yaml:"mapping"`
	// Used when the source of the import should be a file on S3.
	//
	// The file must be either a JSON or a YAML document.
	S3Source *S3Source `field:"optional" json:"s3Source" yaml:"s3Source"`
}

