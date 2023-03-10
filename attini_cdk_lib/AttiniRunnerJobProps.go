// Attini CDK Constructs
package attini_cdk_lib


type AttiniRunnerJobProps struct {
	// A list of shell commands that will be executed by the Runner.
	Commands *[]*string `field:"required" json:"commands" yaml:"commands"`
	// Environment variables that will be set in the shell for the runner job.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// A reference to an {@link AttiniRunner} to use for executing the job. Use {@link AttiniRunner.runnerName} to get a reference.
	//
	// If omitted the Attini default runner will be used.
	Runner *string `field:"optional" json:"runner" yaml:"runner"`
}

