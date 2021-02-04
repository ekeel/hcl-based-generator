package model

import (
	"hclgen/model/jenkins"

	"github.com/hashicorp/hcl2/hcl"
)

// Project is a model to hold all of the parsed configuration information.
type Project struct {
	Seeds []*jenkins.Seed `hcl:"seed,block"`

	Options hcl.Body `hcl:",remain"`
}
