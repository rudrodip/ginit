package cli

type ProjectType struct {
	Name  string
	Alias int
}

var ProjectTypeOptions = map[ProjectType]GenFunc{
	{Name: "CLI", Alias: 1}:     generateCLIProject,
	{Name: "Backend", Alias: 2}: generateBackendProject,
}
