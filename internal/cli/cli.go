package cli

import "fmt"

func Execute() {
	fprojectName := getUserInput("Enter the project name: ")
	moduleName := getUserInput("Enter the module initialize name (e.g., github.com/username/reponame): ")
	projectType := getProjectTypeSelection()

	project := Project{
		Name:   fprojectName,
		Module: moduleName,
		Type:   projectType,
	}

	selectedFunc := ProjectTypeOptions[projectType]
	err := selectedFunc(project)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
