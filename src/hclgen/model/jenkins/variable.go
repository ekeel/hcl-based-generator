package jenkins

type Variable struct {
	Name  string `hcl:",label"`
	Value string `hcl:"value,attr"`
}
