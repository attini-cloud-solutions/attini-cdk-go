// Attini resources
package attini_cdk_lib


type RunnerConfiguration struct {
	IdleTimeToLive *float64 `field:"optional" json:"idleTimeToLive" yaml:"idleTimeToLive"`
	JobTimeout *float64 `field:"optional" json:"jobTimeout" yaml:"jobTimeout"`
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	MaxConcurrentJobs *float64 `field:"optional" json:"maxConcurrentJobs" yaml:"maxConcurrentJobs"`
}

