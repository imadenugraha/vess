package cmd

import (
	"fmt"
	"os"

	"vess/internal/config"
	"vess/internal/generator"
	"vess/internal/logger"

	"github.com/spf13/cobra"
)

var (
	envFile    string
	outputFile string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a Dockerfile for PHP",
	Long: `Generate a Dockerfile based on PHP extensions specified in an env file.
	
The env file should contain PHP extensions in the format:
  PHP_EXTENSIONS=mysqli,pdo_mysql,gd,redis,opcache
  
The generated Dockerfile will include all necessary system dependencies
and installation commands for the specified OS and PHP version.`,
	Example: `  vess generate --os alpine --php-version 8.2 --env-file app.env --output Dockerfile
  vess generate -o ubuntu -p 8.3 -e config.env -f Dockerfile.ubuntu`,
	RunE: runGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&envFile, "env-file", "e", ".env", "Path to env file containing PHP extensions")
	generateCmd.Flags().StringVarP(&outputFile, "output", "f", "Dockerfile", "Output path for generated Dockerfile")
	generateCmd.MarkFlagRequired("env-file")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	log := logger.New(IsVerbose())

	log.Info("Starting Dockerfile generation")
	log.Debug("OS: %s, PHP Version: %s", GetOSType(), GetPHPVersion())
	log.Debug("Env file: %s, Output: %s", envFile, outputFile)

	// Parse env file
	log.Info("Parsing configuration file...")
	cfg, err := config.ParseEnvFile(envFile)
	if err != nil {
		return fmt.Errorf("failed to parse env file: %w", err)
	}

	// Validate configuration
	log.Info("Validating extensions...")
	validator := config.NewValidator()
	if err := validator.Validate(cfg, GetOSType(), GetPHPVersion()); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Generate Dockerfile
	log.Info("Generating Dockerfile...")
	gen := generator.New(GetOSType(), GetPHPVersion())
	content, err := gen.Generate(cfg.Extensions)
	if err != nil {
		return fmt.Errorf("failed to generate Dockerfile: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write Dockerfile: %w", err)
	}

	log.Success("Dockerfile generated successfully: %s", outputFile)
	log.Info("Extensions included: %d", len(cfg.Extensions))

	return nil
}
