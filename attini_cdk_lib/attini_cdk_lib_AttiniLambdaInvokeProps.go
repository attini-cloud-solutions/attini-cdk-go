// Attini CDK Constructs
package attini_cdk_lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
)

type AttiniLambdaInvokeProps struct {
	// Lambda function to invoke.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// Up to 3583 bytes of base64-encoded data about the invoking client to pass to the function.
	ClientContext *string `field:"optional" json:"clientContext" yaml:"clientContext"`
	// Invocation type of the Lambda function.
	InvocationType awsstepfunctionstasks.LambdaInvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The JSON that will be supplied as input to the Lambda function.
	//
	// If not specified then the entire payload will be passed.
	Payload *map[string]interface{} `field:"optional" json:"payload" yaml:"payload"`
	// Version or alias to invoke a published version of the function.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

