package config

import (
	"fmt"
	"strings"

	"vess/internal/extensions"
)

// Validator validates configuration
type Validator struct{}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{}
}

// Validate validates the configuration
func (v *Validator) Validate(cfg *extensions.Config, osType, phpVersion string) error {
	if len(cfg.Extensions) == 0 {
		return fmt.Errorf("no extensions specified")
	}

	// Validate OS
	if osType != "alpine" && osType != "ubuntu" {
		return fmt.Errorf("unsupported OS: %s (must be 'alpine' or 'ubuntu')", osType)
	}

	// Validate PHP version
	validVersions := []string{"7.4", "8.0", "8.1", "8.2", "8.3"}
	if !contains(validVersions, phpVersion) {
		return fmt.Errorf("unsupported PHP version: %s (must be one of: %s)", 
			phpVersion, strings.Join(validVersions, ", "))
	}

	// Validate each extension
	for _, extName := range cfg.Extensions {
		if err := v.validateExtension(extName, osType, phpVersion); err != nil {
			return err
		}
	}

	// Check for conflicts
	if err := v.checkConflicts(cfg.Extensions); err != nil {
		return err
	}

	return nil
}

// validateExtension validates a single extension
func (v *Validator) validateExtension(extName, osType, phpVersion string) error {
	ext, exists := extensions.GetExtension(extName)
	if !exists {
		return &extensions.ValidationError{
			Extension: extName,
			Message:   fmt.Sprintf("unknown extension: %s", extName),
		}
	}

	// Check PHP version support
	if !extensions.SupportsVersion(extName, phpVersion) {
		return &extensions.ValidationError{
			Extension: extName,
			Message: fmt.Sprintf("extension '%s' does not support PHP %s (supported: %s)",
				extName, phpVersion, strings.Join(ext.PHPVersions, ", ")),
		}
	}

	// Check OS support
	if !extensions.SupportsOS(extName, osType) {
		return &extensions.ValidationError{
			Extension: extName,
			Message:   fmt.Sprintf("extension '%s' does not support OS: %s", extName, osType),
		}
	}

	return nil
}

// checkConflicts checks for conflicting extensions
func (v *Validator) checkConflicts(extNames []string) error {
	for _, extName := range extNames {
		ext, exists := extensions.GetExtension(extName)
		if !exists {
			continue
		}

		for _, conflict := range ext.Conflicts {
			if contains(extNames, conflict) {
				return fmt.Errorf("extension '%s' conflicts with '%s'", extName, conflict)
			}
		}
	}
	return nil
}

// contains checks if a slice contains a string
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
