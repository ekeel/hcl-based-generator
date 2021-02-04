package jenkins

type SeedJobEnvironment struct {
	Name      string      `hcl:",label"`
	Variables []*Variable `hcl:"variable,block"`
}
