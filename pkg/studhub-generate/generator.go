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

	if err := os.MkdirAll(serviceDir, 0750); err != nil {
		return err
	}

	serviceRoot, err := os.OpenRoot(serviceDir)
	if err != nil {
		return err
	}
	defer serviceRoot.Close()

	// 1. Define the internal directory tree
	dirs := []string{
		filepath.Join("cmd", "service"),
		filepath.Join(".github", "workflows"),
	}

	for _, dir := range dirs {
		if err := serviceRoot.MkdirAll(dir, 0750); err != nil {
			return err
		}
	}

	// 2. Map template sources to destination paths
	files := map[string]string{
		"templates/main.go.tmpl":            filepath.Join("cmd", "service", "main.go"),
		"templates/ci-cd.yaml.tmpl":         filepath.Join(".github", "workflows", "ci-cd.yaml"),
		"templates/docker_file.tmpl":        "Dockerfile",
		"templates/dockerignore.tmpl":       ".dockerignore",
		"templates/docker-compose.yml.tmpl": "docker-compose.yml",
		"templates/gitignore.tmpl":          ".gitignore",
		"templates/makefile.tmpl":           "Makefile",
		"templates/README.md.tmpl":          "README.md",
	}

	// 3. Render each file
	for tmplSource, destPath := range files {
		if err := renderTemplate(serviceRoot, tmplSource, destPath, data); err != nil {
			return err
		}
	}

	// 4. Automatically initialize Go dependencies and format code
	// initialize a new Go module
	modName := fmt.Sprintf("github.com/MorningBlossom/%s", serviceName)
	// #nosec G204 -- serviceName is passed as a single argument to a fixed executable; it is not interpreted by a shell.
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
