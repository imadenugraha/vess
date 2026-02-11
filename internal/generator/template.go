package generator

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

	"vess/internal/extensions"
)

//go:embed templates/*.tmpl
var templatesFS embed.FS

// TemplateEngine wraps Go's template engine
type TemplateEngine struct {
	templates *template.Template
}

// NewTemplateEngine creates a new template engine
func NewTemplateEngine() (*TemplateEngine, error) {
	funcMap := template.FuncMap{
		"minus1": func(n int) int { return n - 1 },
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseFS(templatesFS, "templates/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	return &TemplateEngine{
		templates: tmpl,
	}, nil
}

// Render renders a template with the given data
func (te *TemplateEngine) Render(templateName string, data interface{}) (string, error) {
	var buf bytes.Buffer

	if err := te.templates.ExecuteTemplate(&buf, templateName, data); err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return buf.String(), nil
}

// TemplateData holds data for template rendering
type TemplateData struct {
	PHPVersion     string
	OSType         string
	BaseImage      string
	BuildDeps      []string
	RuntimeDeps    []string
	Extensions     []*ExtensionData
	HasBuildDeps   bool
	HasRuntimeDeps bool
}

// ExtensionData holds extension-specific data for templates
type ExtensionData struct {
	Name        string
	InstallCmd  string
	PECLInstall bool
}

// PrepareTemplateData prepares data for template rendering
func PrepareTemplateData(osType, phpVersion string, extNames []string) (*TemplateData, error) {
	data := &TemplateData{
		PHPVersion: phpVersion,
		OSType:     osType,
		Extensions: make([]*ExtensionData, 0, len(extNames)),
	}

	// Set base image
	if osType == "alpine" {
		data.BaseImage = fmt.Sprintf("php:%s-fpm-alpine", phpVersion)
		data.BuildDeps = extensions.GetAlpineBuildDeps(extNames)
		data.RuntimeDeps = extensions.GetAlpineRuntimeDeps(extNames)
	} else if osType == "ubuntu" {
		data.BaseImage = fmt.Sprintf("php:%s-fpm", phpVersion)
		data.BuildDeps = extensions.GetUbuntuBuildDeps(extNames)
		data.RuntimeDeps = extensions.GetUbuntuRuntimeDeps(extNames)
	}

	data.HasBuildDeps = len(data.BuildDeps) > 0
	data.HasRuntimeDeps = len(data.RuntimeDeps) > 0

	// Prepare extension data
	for _, extName := range extNames {
		ext, exists := extensions.GetExtension(extName)
		if !exists {
			continue
		}

		osSupport := ext.OSSupport[osType]
		if osSupport == nil {
			continue
		}

		data.Extensions = append(data.Extensions, &ExtensionData{
			Name:        extName,
			InstallCmd:  osSupport.InstallCmd,
			PECLInstall: osSupport.PECLInstall,
		})
	}

	return data, nil
}
