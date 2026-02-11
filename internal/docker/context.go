package docker

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// BuildContext creates a build context for Docker
type BuildContext struct {
	dockerfilePath string
}

// NewBuildContext creates a new build context
func NewBuildContext(dockerfilePath string) *BuildContext {
	return &BuildContext{
		dockerfilePath: dockerfilePath,
	}
}

// CreateTar creates a tar archive containing the Dockerfile
func (bc *BuildContext) CreateTar() (io.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Read Dockerfile
	dockerfileContent, err := os.ReadFile(bc.dockerfilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read Dockerfile: %w", err)
	}

	// Add Dockerfile to tar
	header := &tar.Header{
		Name: "Dockerfile",
		Mode: 0644,
		Size: int64(len(dockerfileContent)),
	}

	if err := tw.WriteHeader(header); err != nil {
		return nil, fmt.Errorf("failed to write tar header: %w", err)
	}

	if _, err := tw.Write(dockerfileContent); err != nil {
		return nil, fmt.Errorf("failed to write to tar: %w", err)
	}

	// Check if there's a directory with additional files
	dir := filepath.Dir(bc.dockerfilePath)
	if dir != "." && dir != "/" {
		// Add other files from the directory if needed
		// For now, we just include the Dockerfile
	}

	return buf, nil
}

// CreateContextFromDirectory creates a tar archive from a directory
func CreateContextFromDirectory(dir string) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Create tar header
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		// Update name to be relative to the directory
		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		// Write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// Write file content
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(tw, file)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create context: %w", err)
	}

	return buf, nil
}
