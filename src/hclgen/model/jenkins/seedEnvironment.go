package jenkins

type SeedEnvironment struct {
	Name      string      `hcl:",label"`
	Variables []*Variable `hcl:"variable,block"`
}
