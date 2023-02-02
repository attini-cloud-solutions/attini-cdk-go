// Attini resources
package attini_cdk_lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

type DeploymentPlanProps struct {
	Definition awsstepfunctions.Chain `field:"required" json:"definition" yaml:"definition"`
}

