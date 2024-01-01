package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func isValidSelection(selection string, maxOptions int) bool {
	if selection == "" {
		return false
	}

	num := int(selection[0] - '0')
	return num >= 1 && num <= maxOptions
}

func getProjectTypeSelection() ProjectType {
	fmt.Println("Choose project type:")
	for pt := range ProjectTypeOptions {
		fmt.Printf("%d. %s\n", pt.Alias, pt.Name)
	}

	var selection string
	for {
		selection = getUserInput("Enter the number corresponding to your choice:")
		if isValidSelection(selection, len(ProjectTypeOptions)) {
			break
		}
		fmt.Println("Invalid selection. Please enter a valid number.")
	}

	index, _ := strconv.Atoi(selection)

	for pt := range ProjectTypeOptions {
		if pt.Alias == index {
			return pt
		}
	}

	return ProjectType{}
}

type Folder struct {
	Name     string
	Children []interface{}
}

type File struct {
	Name    string
	Content string
}

func generateMyCLIStructure(project Project) Folder {
	rootFolderName := project.Name

	fileStructure := Folder{
		Name: rootFolderName,
		Children: []interface{}{
			Folder{
				Name: "cmd",
				Children: []interface{}{
					Folder{
						Name: rootFolderName,
						Children: []interface{}{
							File{
								Name:    "main.go",
								Content: `package main`,
							},
						},
					},
				},
			},
			Folder{
				Name: "internal",
				Children: []interface{}{
					Folder{
						Name: "app",
						Children: []interface{}{
							File{
								Name:    "app.go",
								Content: `package app`,
							},
							File{
								Name:    "config.go",
								Content: `package app`,
							},
							Folder{
								Name: "commands",
							},
						},
					},
				},
			},
			Folder{
				Name: "pkg",
			},
			File{
				Name:    ".gitignore",
				Content: `bin`,
			},
			File{
				Name:    "README.md",
				Content: `# project documentation`,
			},
		},
	}

	return fileStructure
}

func generateBackendProjectStructure(project Project) Folder {
	rootFolderName := project.Name

	fileStructure := Folder{
		Name: rootFolderName,
		Children: []interface{}{
			Folder{
				Name: "cmd",
				Children: []interface{}{
					Folder{
						Name: rootFolderName,
						Children: []interface{}{
							File{
								Name:    "main.go",
								Content: `package main`,
							},
						},
					},
				},
			},
			Folder{
				Name: "internal",
				Children: []interface{}{
					Folder{
						Name: "app",
						Children: []interface{}{
							File{
								Name:    "app.go",
								Content: `package app`,
							},
							File{
								Name:    "config.go",
								Content: `package app`,
							},
							Folder{
								Name: "handlers",
							},
							Folder{
								Name: "models",
							},
							Folder{
								Name: "repositories",
							},
							Folder{
								Name: "services",
							},
						},
					},
				},
			},
			Folder{
				Name: "pkg",
			},
			File{
				Name:    ".gitignore",
				Content: `bin`,
			},
			File{
				Name:    "README.md",
				Content: `# project documentation`,
			},
		},
	}

	return fileStructure
}

func printFolderStructure(node Folder, indent string) {
	fmt.Println(indent + "|-- " + node.Name)
	for _, child := range node.Children {
		switch v := child.(type) {
		case Folder:
			printFolderStructure(v, indent+"|   ")
		case File:
			fmt.Println(indent + "|   |-- " + v.Name)
		}
	}
}
