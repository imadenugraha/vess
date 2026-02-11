package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

// Client wraps the Docker SDK client
type Client struct {
	cli *client.Client
	ctx context.Context
}

// NewClient creates a new Docker client
func NewClient() (*Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}

	return &Client{
		cli: cli,
		ctx: context.Background(),
	}, nil
}

// Close closes the Docker client
func (c *Client) Close() error {
	if c.cli != nil {
		return c.cli.Close()
	}
	return nil
}

// GetClient returns the underlying Docker client
func (c *Client) GetClient() *client.Client {
	return c.cli
}

// GetContext returns the context
func (c *Client) GetContext() context.Context {
	return c.ctx
}

// Ping checks if Docker daemon is available
func (c *Client) Ping() error {
	_, err := c.cli.Ping(c.ctx)
	if err != nil {
		return fmt.Errorf("failed to ping Docker daemon: %w", err)
	}
	return nil
}
