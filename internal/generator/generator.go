package generator

import (
	"fmt"
)

// Generator generates Dockerfiles
type Generator struct {
	osType     string
	phpVersion string
	engine     *TemplateEngine
}

// New creates a new Dockerfile generator
func New(osType, phpVersion string) *Generator {
	engine, _ := NewTemplateEngine()
	return &Generator{
		osType:     osType,
		phpVersion: phpVersion,
		engine:     engine,
	}
}

// Generate generates a Dockerfile from extensions
func (g *Generator) Generate(extNames []string) (string, error) {
	// Prepare template data
	data, err := PrepareTemplateData(g.osType, g.phpVersion, extNames)
	if err != nil {
		return "", fmt.Errorf("failed to prepare template data: %w", err)
	}

	// Select template based on OS
	var templateName string
	if g.osType == "alpine" {
		templateName = "alpine.dockerfile.tmpl"
	} else if g.osType == "ubuntu" {
		templateName = "ubuntu.dockerfile.tmpl"
	} else {
		return "", fmt.Errorf("unsupported OS: %s", g.osType)
	}

	// Render template
	content, err := g.engine.Render(templateName, data)
	if err != nil {
		return "", fmt.Errorf("failed to render template: %w", err)
	}

	return content, nil
}
