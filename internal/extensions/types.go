package extensions

// Extension represents a PHP extension with all its metadata
type Extension struct {
	Name          string                `json:"name"`
	Description   string                `json:"description"`
	PHPVersions   []string              `json:"php_versions"`   // Supported PHP versions
	OSSupport     map[string]*OSSupport `json:"os_support"`     // OS-specific installation info
	Conflicts     []string              `json:"conflicts"`      // Conflicting extensions
	ConfigureArgs []string              `json:"configure_args"` // Additional configure arguments
}

// OSSupport contains OS-specific installation information
type OSSupport struct {
	BuildDeps   []string `json:"build_deps"`   // Build-time dependencies
	RuntimeDeps []string `json:"runtime_deps"` // Runtime dependencies
	InstallCmd  string   `json:"install_cmd"`  // Installation command
	PECLInstall bool     `json:"pecl_install"` // Whether to use PECL
}

// Config represents the parsed configuration
type Config struct {
	Extensions []string          `json:"extensions"`
	Metadata   map[string]string `json:"metadata"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Extension string
	Message   string
}

func (e *ValidationError) Error() string {
	return e.Message
}
