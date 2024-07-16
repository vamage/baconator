package main

import (
	"fmt"
	"github.com/forensicanalysis/gitfs"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type Config struct {
	Variables []*Variable `hcl:"variable,block"`
}

type Variable struct {
	Test *string `hcl:"test"`
}

func main() {
	fsys, _ := gitfs.New("https://github.com/vamage/baconator-modules")
	wrap := tfconfig.WrapFS(fsys)
	module, _ := tfconfig.LoadModuleFromFilesystem(wrap, "testing-variables")
	fmt.Printf("module %+v\n", module.Variables["booo"])
	fmt.Printf("%+V\n", module.Diagnostics)
}
