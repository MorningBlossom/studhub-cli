package generator

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

type TemplateData struct {
	ServiceName string
}

func renderTemplate(tmplPath string, destPath string, data TemplateData) error {
	content, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return err
	}

	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(content))
	if err != nil {
		return err
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}