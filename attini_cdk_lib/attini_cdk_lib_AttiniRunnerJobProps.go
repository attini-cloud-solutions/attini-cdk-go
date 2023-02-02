// Attini resources
package attini_cdk_lib


type AttiniRunnerJobProps struct {
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	Runner *string `field:"optional" json:"runner" yaml:"runner"`
}

