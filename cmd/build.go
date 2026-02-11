package cmd

import (
	"fmt"

	"vess/internal/docker"
	"vess/internal/logger"

	"github.com/spf13/cobra"
)

var (
	dockerfile string
	tag        string
	noCache    bool
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a Docker image from Dockerfile",
	Long: `Build a Docker image from a generated Dockerfile.
	
This command uses the Docker SDK to build the image and streams
the build output to the terminal. You can specify a custom tag
and control caching behavior.`,
	Example: `  vess build --dockerfile Dockerfile --tag my-php:8.2
  vess build -d Dockerfile.alpine -t my-app:latest --no-cache`,
	RunE: runBuild,
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVarP(&dockerfile, "dockerfile", "d", "Dockerfile", "Path to Dockerfile")
	buildCmd.Flags().StringVarP(&tag, "tag", "t", "", "Image tag (e.g., my-php:8.2)")
	buildCmd.Flags().BoolVar(&noCache, "no-cache", false, "Do not use cache when building")
	buildCmd.MarkFlagRequired("tag")
}

func runBuild(cmd *cobra.Command, args []string) error {
	log := logger.New(IsVerbose())
	
	log.Info("Starting Docker image build")
	log.Debug("Dockerfile: %s, Tag: %s, NoCache: %v", dockerfile, tag, noCache)

	// Create Docker client
	client, err := docker.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer client.Close()

	// Build image
	log.Info("Building image: %s", tag)
	builder := docker.NewBuilder(client, log)
	if err := builder.Build(dockerfile, tag, noCache); err != nil {
		return fmt.Errorf("failed to build image: %w", err)
	}

	log.Success("Image built successfully: %s", tag)
	log.Info("Run: docker run --rm %s php -m", tag)
	
	return nil
}
