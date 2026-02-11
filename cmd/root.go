package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	osType     string
	phpVersion string
	verbose    bool
)

var rootCmd = &cobra.Command{
	Use:   "vess",
	Short: "PHP Dockerfile generator and builder",
	Long: `vess is a CLI tool that generates OS-specific PHP Dockerfiles,
builds Docker images, and exports PHP extension metadata to JSON.

Supports PHP 7.4+ with Alpine and Ubuntu base images.`,
	Version: "1.0.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags available to all subcommands
	rootCmd.PersistentFlags().StringVarP(&osType, "os", "o", "alpine", "Operating system (alpine, ubuntu)")
	rootCmd.PersistentFlags().StringVarP(&phpVersion, "php-version", "p", "8.3", "PHP version (7.4, 8.0, 8.1, 8.2, 8.3)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

// GetOSType returns the configured OS type
func GetOSType() string {
	return osType
}

// GetPHPVersion returns the configured PHP version
func GetPHPVersion() string {
	return phpVersion
}

// IsVerbose returns whether verbose mode is enabled
func IsVerbose() bool {
	return verbose
}

// PrintError prints an error message and exits
func PrintError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
}
