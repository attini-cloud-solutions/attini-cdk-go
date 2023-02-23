// Attini CDK Constructs
package attini_cdk_lib


// Startup instructions for the Attini Runner.
type Startup struct {
	// List of shell commands that are executed when the Attini Runner starts.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// The number of seconds the startup commands can execute before the Attini Runner aborts the execution.
	CommandsTimeout *float64 `field:"optional" json:"commandsTimeout" yaml:"commandsTimeout"`
}

