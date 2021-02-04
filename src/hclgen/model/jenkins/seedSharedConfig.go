package jenkins

type SeedSharedConfig struct {
	Variables []*Variable `hcl:"variable,block"`
}
