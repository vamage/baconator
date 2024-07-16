package terraform

import (
	"baconator/api"
	"fmt"
	"github.com/forensicanalysis/gitfs"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func ReadTF(url, moduleName string) (resource *api.Resource, err error) {
	resource = &api.Resource{
		Name:           "moduleName",
		ResourceInputs: nil,
	}
	fsys, err := gitfs.New(url)
	if err != nil {
		return nil, err
	}
	wrap := tfconfig.WrapFS(fsys)
	module, diag := tfconfig.LoadModuleFromFilesystem(wrap, moduleName)
	if diag.Err() != nil {
		return nil, fmt.Errorf("error reading module %s, %v", moduleName, err)
	}
	for n, v := range module.Variables {
		i := api.Input{
			Name:        n,
			Description: v.Description,
			Type: api.OptString{
				Value: v.Type,
				Set:   true,
			},
		}
		resource.ResourceInputs = append(resource.ResourceInputs, i)

	}
	return resource, nil
}
