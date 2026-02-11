package exporter

// ExportData represents the exported extension metadata
type ExportData struct {
	OS         string              `json:"os"`
	PHPVersion string              `json:"php_version"`
	Extensions []*ExtensionExport  `json:"extensions"`
}

// AllExportData represents exported data for all OS/PHP combinations
type AllExportData struct {
	Data map[string]map[string]*ExportData `json:"data"`
}

// ExtensionExport represents a single extension's export data
type ExtensionExport struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	BuildDeps     []string `json:"build_dependencies"`
	RuntimeDeps   []string `json:"runtime_dependencies"`
	InstallCmd    string   `json:"install_command"`
	PECLInstall   bool     `json:"pecl_install"`
	PHPVersions   []string `json:"supported_php_versions"`
	Conflicts     []string `json:"conflicts"`
	ConfigureArgs []string `json:"configure_args,omitempty"`
}

// FormatSummary creates a summary of the export
type FormatSummary struct {
	TotalExtensions int      `json:"total_extensions"`
	PECLExtensions  int      `json:"pecl_extensions"`
	CoreExtensions  int      `json:"core_extensions"`
	ExtensionNames  []string `json:"extension_names"`
}

// CreateSummary creates a summary from export data
func CreateSummary(data *ExportData) *FormatSummary {
	summary := &FormatSummary{
		TotalExtensions: len(data.Extensions),
		ExtensionNames:  make([]string, 0, len(data.Extensions)),
	}

	for _, ext := range data.Extensions {
		summary.ExtensionNames = append(summary.ExtensionNames, ext.Name)
		if ext.PECLInstall {
			summary.PECLExtensions++
		} else {
			summary.CoreExtensions++
		}
	}

	return summary
}
