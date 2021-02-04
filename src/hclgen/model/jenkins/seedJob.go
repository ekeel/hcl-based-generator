package jenkins

type SeedJob struct {
	Name         string                `hcl:",label"`
	Description  string                `hcl:"description,optional"`
	ScriptPath   string                `hcl:"script_path,attr"`
	Environments []*SeedJobEnvironment `hcl:"env,block"`
	View         *SeedJobView          `hcl:"view,block"`
	Variables    []*Variable           `hcl:"variable,block"`
}
