// Attini CDK Constructs
package attini_cdk_lib


type AttiniCdkProps struct {
	// The path to the CDK project.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Passed to the CDK --app option.
	App *string `field:"optional" json:"app" yaml:"app"`
	// Passed to the CDK --build option.
	BuildCommands *string `field:"optional" json:"buildCommands" yaml:"buildCommands"`
	// Passed to the CDK --build-exclude option.
	BuildExclude *[]*string `field:"optional" json:"buildExclude" yaml:"buildExclude"`
	// Passed to the CDK --context option.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Configure if you want Attini to perform a diff check on the CDK app before any changes are applied.
	//
	// If there are any changes performed on the CDK app, then manual approval will be required before the changes
	// are applied.
	Diff *DiffProps `field:"optional" json:"diff" yaml:"diff"`
	// Environment variables that will be set in the shell for the runner job.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Passed to the CDK --force option.
	Force *bool `field:"optional" json:"force" yaml:"force"`
	// Passed to the CDK --notification-arns option.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Passed to the CDK --plugins option.
	Plugins *[]*string `field:"optional" json:"plugins" yaml:"plugins"`
	// Passed to the CDK --role-arn option.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A reference to an {@link AttiniRunner} to use for executing the job. Use {@link AttiniRunner.runnerName} to get a reference.
	//
	// If omitted the Attini default runner will be used.
	Runner *string `field:"optional" json:"runner" yaml:"runner"`
	// Passed to the CDK --parameters options.
	StackConfiguration *[]*StackConfigurationProps `field:"optional" json:"stackConfiguration" yaml:"stackConfiguration"`
	// Stacks to deploy.
	Stacks *[]*string `field:"optional" json:"stacks" yaml:"stacks"`
}

