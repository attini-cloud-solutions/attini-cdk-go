// Attini resources
package attini_cdk_lib


type Project struct {
	Path *string `field:"required" json:"path" yaml:"path"`
	BuildDir *string `field:"optional" json:"buildDir" yaml:"buildDir"`
	Template *string `field:"optional" json:"template" yaml:"template"`
}

