// Package: main
// This package is for testing terraform only
package main

import (
	"fmt"

	"github.com/forensicanalysis/gitfs"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

// Config is a struct that contains the configuration for the terraform module.
type Config struct {
	Variables []*Variable `hcl:"variable,block"`
}

// Variable is a struct that contains the variables for the terraform module.
type Variable struct {
	Test *string `hcl:"test"`
}

func main() {
	fsys, _ := gitfs.New("https://github.com/vamage/baconator-modules")
	wrap := tfconfig.WrapFS(fsys)
	module, _ := tfconfig.LoadModuleFromFilesystem(wrap, "testing-variables")
	fmt.Printf("module %+v\n", module.Variables["booo"])
	fmt.Printf("%+v\n", module.Diagnostics)
}
