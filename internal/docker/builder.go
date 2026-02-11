package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
)

// Logger interface for builder output
type Logger interface {
	Info(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Error(format string, args ...interface{})
}

// Builder builds Docker images
type Builder struct {
	client *Client
	logger Logger
}

// NewBuilder creates a new Docker builder
func NewBuilder(client *Client, logger Logger) *Builder {
	return &Builder{
		client: client,
		logger: logger,
	}
}

// Build builds a Docker image from a Dockerfile
func (b *Builder) Build(dockerfilePath, tag string, noCache bool) error {
	// Check if Docker daemon is available
	if err := b.client.Ping(); err != nil {
		return err
	}

	// Create build context
	b.logger.Debug("Creating build context...")
	ctx := NewBuildContext(dockerfilePath)
	buildContext, err := ctx.CreateTar()
	if err != nil {
		return fmt.Errorf("failed to create build context: %w", err)
	}

	// Build options
	buildOptions := types.ImageBuildOptions{
		Tags:       []string{tag},
		Dockerfile: "Dockerfile",
		Remove:     true,
		NoCache:    noCache,
	}

	// Build image
	b.logger.Debug("Starting Docker build...")
	resp, err := b.client.GetClient().ImageBuild(
		b.client.GetContext(),
		buildContext,
		buildOptions,
	)
	if err != nil {
		return fmt.Errorf("failed to build image: %w", err)
	}
	defer resp.Body.Close()

	// Stream build output
	if err := b.streamOutput(resp.Body); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	return nil
}

// streamOutput streams Docker build output
func (b *Builder) streamOutput(reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	for {
		var message struct {
			Stream      string `json:"stream"`
			Error       string `json:"error"`
			ErrorDetail struct {
				Message string `json:"message"`
			} `json:"errorDetail"`
		}

		if err := decoder.Decode(&message); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if message.Error != "" {
			return fmt.Errorf("build error: %s", message.Error)
		}

		if message.Stream != "" {
			fmt.Fprint(os.Stdout, message.Stream)
		}
	}

	return nil
}

// ListImages lists Docker images
func (b *Builder) ListImages() ([]image.Summary, error) {
	images, err := b.client.GetClient().ImageList(
		b.client.GetContext(),
		image.ListOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list images: %w", err)
	}
	return images, nil
}

// RemoveImage removes a Docker image
func (b *Builder) RemoveImage(imageID string, force bool) error {
	_, err := b.client.GetClient().ImageRemove(
		b.client.GetContext(),
		imageID,
		image.RemoveOptions{Force: force},
	)
	if err != nil {
		return fmt.Errorf("failed to remove image: %w", err)
	}
	return nil
}
