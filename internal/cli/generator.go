package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type GenFunc func(project Project) error

type Project struct {
	Name   string
	Module string
	Type   ProjectType
}

func generateCLIProject(project Project) error {
	fmt.Println("Generating CLI project...")
	fileStructure := generateMyCLIStructure(project)
	rootPath := filepath.Join(".")

	if err := createFolderStructure(rootPath, fileStructure); err != nil {
		return err
	}

	if err := os.Chdir(filepath.Join(".", project.Name)); err != nil {
		return err
	}

	if err := initialize(project.Module); err != nil {
		return err
	}

	printFolderStructure(fileStructure, "")
	fmt.Printf("Project %s successfully generated!\n", project.Name)

	return nil
}

func generateBackendProject(project Project) error {
	fmt.Println("Generating Backend project...")

	fileStructure := generateBackendProjectStructure(project)
	rootPath := filepath.Join(".")
	fmt.Println(rootPath)

	if err := createFolderStructure(rootPath, fileStructure); err != nil {
		return err
	}

	if err := os.Chdir(filepath.Join(".", project.Name)); err != nil {
		return err
	}

	if err := initialize(project.Module); err != nil {
		return err
	}

	printFolderStructure(fileStructure, "")
	fmt.Printf("Project %s successfully generated!\n", project.Name)

	return nil
}

func generateBlankProject(project Project) error {
	fmt.Println("Generating Backend project...")

	fileStructure := generateBlankProjectStructure(project)
	rootPath := filepath.Join(".")
	fmt.Println(rootPath)

	if err := createFolderStructure(rootPath, fileStructure); err != nil {
		return err
	}

	if err := os.Chdir(filepath.Join(".", project.Name)); err != nil {
		return err
	}

	if err := initialize(project.Module); err != nil {
		return err
	}

	printFolderStructure(fileStructure, "")
	fmt.Printf("Project %s successfully generated!\n", project.Name)

	return nil
}

func createFolderStructure(rootPath string, node Folder) error {
	rootFullPath := filepath.Join(rootPath, node.Name)

	if err := os.Mkdir(rootFullPath, 0755); err != nil {
		return err
	}

	for _, child := range node.Children {
		switch v := child.(type) {
		case Folder:
			if err := createFolderStructure(rootFullPath, v); err != nil {
				return err
			}
		case File:
			filePath := filepath.Join(rootFullPath, v.Name)
			if err := os.WriteFile(filePath, []byte(v.Content), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}

func initialize(moduleName string) error {
	if err := goModInit(moduleName); err != nil {
		return err
	}
	if err := gitInit(); err != nil {
		return err
	}

	return nil
}

func goModInit(moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func gitInit() error {
	cmd := exec.Command("git", "init")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
