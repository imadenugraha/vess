package exporter

import (
	"encoding/json"
	"fmt"

	"vess/internal/extensions"
)

// Exporter exports extension metadata
type Exporter struct{}

// New creates a new exporter
func New() *Exporter {
	return &Exporter{}
}

// Export exports extension metadata to JSON
func (e *Exporter) Export(osType, phpVersion string) ([]byte, error) {
	// Get all extensions
	registry := extensions.GetRegistry()

	// Filter and format extensions
	data := &ExportData{
		OS:         osType,
		PHPVersion: phpVersion,
		Extensions: make([]*ExtensionExport, 0),
	}

	for name, ext := range registry {
		// Check if extension supports this version and OS
		if !extensions.SupportsVersion(name, phpVersion) {
			continue
		}
		if !extensions.SupportsOS(name, osType) {
			continue
		}

		// Get OS-specific support
		osSupport := ext.OSSupport[osType]
		if osSupport == nil {
			continue
		}

		exportExt := &ExtensionExport{
			Name:          ext.Name,
			Description:   ext.Description,
			BuildDeps:     osSupport.BuildDeps,
			RuntimeDeps:   osSupport.RuntimeDeps,
			InstallCmd:    osSupport.InstallCmd,
			PECLInstall:   osSupport.PECLInstall,
			PHPVersions:   ext.PHPVersions,
			Conflicts:     ext.Conflicts,
			ConfigureArgs: ext.ConfigureArgs,
		}

		data.Extensions = append(data.Extensions, exportExt)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return jsonData, nil
}

// ExportAll exports all extensions for all OS and PHP version combinations
func (e *Exporter) ExportAll() ([]byte, error) {
	allData := &AllExportData{
		Data: make(map[string]map[string]*ExportData),
	}

	osTypes := []string{"alpine", "ubuntu"}
	phpVersions := []string{"7.4", "8.0", "8.1", "8.2", "8.3"}

	for _, osType := range osTypes {
		allData.Data[osType] = make(map[string]*ExportData)

		for _, phpVersion := range phpVersions {
			data, err := e.exportForVersion(osType, phpVersion)
			if err != nil {
				return nil, err
			}
			allData.Data[osType][phpVersion] = data
		}
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(allData, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return jsonData, nil
}

// exportForVersion exports data for a specific version (internal helper)
func (e *Exporter) exportForVersion(osType, phpVersion string) (*ExportData, error) {
	registry := extensions.GetRegistry()

	data := &ExportData{
		OS:         osType,
		PHPVersion: phpVersion,
		Extensions: make([]*ExtensionExport, 0),
	}

	for name, ext := range registry {
		if !extensions.SupportsVersion(name, phpVersion) {
			continue
		}
		if !extensions.SupportsOS(name, osType) {
			continue
		}

		osSupport := ext.OSSupport[osType]
		if osSupport == nil {
			continue
		}

		exportExt := &ExtensionExport{
			Name:          ext.Name,
			Description:   ext.Description,
			BuildDeps:     osSupport.BuildDeps,
			RuntimeDeps:   osSupport.RuntimeDeps,
			InstallCmd:    osSupport.InstallCmd,
			PECLInstall:   osSupport.PECLInstall,
			PHPVersions:   ext.PHPVersions,
			Conflicts:     ext.Conflicts,
			ConfigureArgs: ext.ConfigureArgs,
		}

		data.Extensions = append(data.Extensions, exportExt)
	}

	return data, nil
}
