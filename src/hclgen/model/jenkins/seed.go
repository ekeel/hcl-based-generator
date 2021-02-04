package jenkins

import (
	"fmt"
	"strings"
)

type Seed struct {
	Name         string             `hcl:",label"`
	Constants    []*SeedConstant    `hcl:"constant,block"`
	SharedConfig *SeedSharedConfig  `hcl:"shared_config,block"`
	Jobs         []*SeedJob         `hcl:"job,block"`
	Environments []*SeedEnvironment `hcl:"environment,block"`

	ConstantStrings []string
}

// GenerateConstants generates a constant string from each constant block.
func (seed *Seed) GenerateConstants() (err error) {
	for _, constant := range seed.Constants {
		switch strings.ToLower(constant.Type) {
		case "string":
			str := fmt.Sprintf("%v %v = '%v'", constant.Type, constant.Name, constant.Value)
			seed.ConstantStrings = append(seed.ConstantStrings, str)
		default:
			str := fmt.Sprintf("%v %v = %v", constant.Type, constant.Name, constant.Value)
			seed.ConstantStrings = append(seed.ConstantStrings, str)
		}
	}

	return nil
}
