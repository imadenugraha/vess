package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"vess/internal/extensions"
)

// ParseEnvFile parses an env file and returns the configuration
func ParseEnvFile(filepath string) (*extensions.Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	config := &extensions.Config{
		Extensions: []string{},
		Metadata:   make(map[string]string),
	}

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE format
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid syntax at line %d: %s", lineNum, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		value = strings.Trim(value, `"'`)

		// Handle PHP_EXTENSIONS specifically
		if key == "PHP_EXTENSIONS" {
			exts := parseExtensions(value)
			config.Extensions = append(config.Extensions, exts...)
		} else {
			config.Metadata[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if len(config.Extensions) == 0 {
		return nil, fmt.Errorf("no extensions specified in PHP_EXTENSIONS")
	}

	return config, nil
}

// parseExtensions parses comma-separated extension names
func parseExtensions(value string) []string {
	exts := strings.Split(value, ",")
	result := make([]string, 0, len(exts))

	for _, ext := range exts {
		ext = strings.TrimSpace(ext)
		if ext != "" {
			result = append(result, ext)
		}
	}

	return result
}
