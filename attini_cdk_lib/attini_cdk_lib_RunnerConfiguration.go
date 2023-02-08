// Attini CDK Constructs
package attini_cdk_lib


type RunnerConfiguration struct {
	// The number of seconds the Attini Runner will stay alive without any jobs executing.
	//
	// New jobs will reset the countdown.
	IdleTimeToLive *float64 `field:"optional" json:"idleTimeToLive" yaml:"idleTimeToLive"`
	// The number of seconds a job can execute before the Attini Runner aborts the execution.
	JobTimeout *float64 `field:"optional" json:"jobTimeout" yaml:"jobTimeout"`
	// The log level of the Attini Runner.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// The max amount of concurrent jobs the Attini Runner will execute.
	MaxConcurrentJobs *float64 `field:"optional" json:"maxConcurrentJobs" yaml:"maxConcurrentJobs"`
}

