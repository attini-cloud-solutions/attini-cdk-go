// Attini CDK Constructs
package attini_cdk_lib


type Project struct {
	// The path to the SAM App in the Attini project.
	//
	// The Path should always start from the root of the Attini project.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The path to a directory where the built artifacts are stored.
	//
	// Only needed if Attini did not perform the SAM build and a custom build dir was specified with the --build-dir option.
	// For more information, see the {@link https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html AWS SAM CLI build documentation}.
	BuildDir *string `field:"optional" json:"buildDir" yaml:"buildDir"`
	// The path and file name of AWS SAM template file.
	//
	// Only needed if Attini did not perform the SAM build and a custom template name was specified with the --template option.
	// For more information, see the {@link https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html AWS SAM CLI build documentation}.
	Template *string `field:"optional" json:"template" yaml:"template"`
}

