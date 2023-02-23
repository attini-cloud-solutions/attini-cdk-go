// Attini CDK Constructs
package attini_cdk_lib


type S3Source struct {
	// The S3 Bucket where the document is located.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 Key of the document to import.
	Key *string `field:"required" json:"key" yaml:"key"`
}

