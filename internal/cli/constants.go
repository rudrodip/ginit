package cli

type ProjectType struct {
	Name  string
	Alias int
}

var ProjectTypeOptions = map[ProjectType]GenFunc{
	{Name: "Blank", Alias: 1}:   generateBlankProject,
	{Name: "CLI", Alias: 2}:     generateCLIProject,
	{Name: "Backend", Alias: 3}: generateBackendProject,
}
