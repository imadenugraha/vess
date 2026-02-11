package cmd

import (
	"fmt"
	"os"

	"vess/internal/exporter"
	"vess/internal/logger"

	"github.com/spf13/cobra"
)

var (
	exportOutput string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export PHP extension metadata to JSON",
	Long: `Export comprehensive PHP extension metadata to a JSON file.
	
The JSON output includes extension names, system dependencies,
installation commands, build flags, conflicts, and compatibility
information for the specified OS and PHP version.`,
	Example: `  vess export --os alpine --php-version 8.2 --output extensions.json
  vess export -o ubuntu -p 8.3 --output ubuntu-ext.json`,
	RunE: runExport,
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVar(&exportOutput, "output", "extensions.json", "Output path for JSON file")
}

func runExport(cmd *cobra.Command, args []string) error {
	log := logger.New(IsVerbose())
	
	log.Info("Exporting extension metadata")
	log.Debug("OS: %s, PHP Version: %s", GetOSType(), GetPHPVersion())
	log.Debug("Output: %s", exportOutput)

	// Export metadata
	exp := exporter.New()
	data, err := exp.Export(GetOSType(), GetPHPVersion())
	if err != nil {
		return fmt.Errorf("failed to export metadata: %w", err)
	}

	// Write to file
	if err := os.WriteFile(exportOutput, data, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	log.Success("Metadata exported successfully: %s", exportOutput)
	
	return nil
}
