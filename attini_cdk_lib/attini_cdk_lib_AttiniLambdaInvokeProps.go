// Attini resources
package attini_cdk_lib


type AttiniLambdaInvokeProps struct {
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	Payload *map[string]interface{} `field:"optional" json:"payload" yaml:"payload"`
}

