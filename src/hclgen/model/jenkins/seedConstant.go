package jenkins

type SeedConstant struct {
	Name    string `hcl:",label"`
	Type    string `hcl:"type,attr"`
	Value   string `hcl:"value,attr"`
}
