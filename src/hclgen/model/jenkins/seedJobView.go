package jenkins

type SeedJobView struct {
	Name        string `hcl:",label"`
	Description string `hcl:"description,optional"`
}
