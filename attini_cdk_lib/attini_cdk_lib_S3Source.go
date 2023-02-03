// Attini CDK Constructs
package attini_cdk_lib


type S3Source struct {
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	Key *string `field:"required" json:"key" yaml:"key"`
}

