package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GenerateService(serviceName string) error {
	data := TemplateData{ServiceName: serviceName}
	serviceDir := serviceName

	if err := os.MkdirAll(serviceDir, 0755); err != nil {
		return err
	}

	// 1. Define the internal directory tree
	dirs := []string{
		filepath.Join(serviceName, "cmd", "service"),
		filepath.Join(serviceName, ".github", "workflows"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// 2. Map template sources to destination paths
	files := map[string]string{
		"templates/main.go.tmpl":       filepath.Join(serviceName, "cmd", "service", "main.go"),
		"templates/ci-cd.yaml.tmpl":    filepath.Join(serviceName, ".github", "workflows", "ci-cd.yaml"),
		"templates/docker_file.tmpl":   filepath.Join(serviceName, "Dockerfile"),
		"templates/dockerignore.tmpl":  filepath.Join(serviceName, ".dockerignore"),
		"templates/gitignore.tmpl":     filepath.Join(serviceName, ".gitignore"),
		"templates/makefile.tmpl":      filepath.Join(serviceName, "Makefile"),
		"templates/README.md.tmpl":     filepath.Join(serviceName, "README.md"),
	}

	// 3. Render each file
	for tmplSource, destPath := range files {
		if err := renderTemplate(tmplSource, destPath, data); err != nil {
			return err
		}
	}

	// 4. Automatically initialize Go dependencies and format code
	// initialize a new Go module
	modName := fmt.Sprintf("github.com/MorningBlossom/%s", serviceName)
	initCmd := exec.Command("go", "mod", "init", modName)
	initCmd.Dir = serviceDir
	initCmd.Stdout = os.Stdout
	initCmd.Stderr = os.Stderr
	if err := initCmd.Run(); err != nil {
		fmt.Printf("Failed to initialize Go module: %v\n", err)
		return err
	}
	modName1 := exec.Command("go", "get", "github.com/jackc/pgx/v5")
	modName1.Dir = serviceDir
	modName1.Stdout = os.Stdout
	modName1.Stderr = os.Stderr
	if err := modName1.Run(); err != nil {
		fmt.Printf("Failed to get pgx dependency: %v\n", err)
		return err
	}
	// Run `go mod tidy` to resolve any imports in your templates
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = serviceDir
	tidyCmd.Stdout = os.Stdout
	tidyCmd.Stderr = os.Stderr
	if err := tidyCmd.Run(); err != nil {
		return err
	}

	// Run `go fmt` to ensure the generated code is perfectly formatted
	fmtCmd := exec.Command("go", "fmt", "./...")
	fmtCmd.Dir = serviceDir
	fmtCmd.Stdout = os.Stdout
	fmtCmd.Stderr = os.Stderr
	if err := fmtCmd.Run(); err != nil {
		return err
	}

	return nil
}