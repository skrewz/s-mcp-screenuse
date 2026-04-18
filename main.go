package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "wayland-screenshot"}, nil)

	server.AddTool(&mcp.Tool{
		Name:        "take_screenshot",
		Description: "Capture a screenshot of the entire Wayland screen using grim",
		InputSchema: map[string]any{
			"type":       "object",
			"properties": map[string]any{},
		},
	}, takeScreenshot)

	if err := initGrim(); err != nil {
		log.Fatalf("Initialization failed: %v", err)
	}
	log.Println("Wayland screenshot server ready")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}

func initGrim() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "grim", "-")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("grim validation failed: %w, stderr: %s", err, stderr.String())
	}
	return nil
}

func takeScreenshot(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cmd := exec.CommandContext(ctx, "grim", "-")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("grim capture failed: %v, stderr: %s", err, stderr.String())},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.ImageContent{
				Data:     stdout.Bytes(),
				MIMEType: "image/png",
			},
		},
	}, nil
}
