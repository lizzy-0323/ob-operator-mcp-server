package main

import (
	"fmt"
	okctl "ob-operator-mcp-server/pkg/okctl"

	"github.com/mark3labs/mcp-go/server"
)

var DefaultToolSets = []string{"all"}

func main() {
	s := server.NewMCPServer("okctl-mcp-server", "0.0.1")

	toolSets, err := okctl.InitToolsets(DefaultToolSets)
	if err != nil {
		fmt.Printf("Failed to initialize toolsets: %v\n", err)
		return
	}
	toolSets.RegisterTools(s)
	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
