package main

import (
	"fmt"
	"ob-operator-mcp-server/pkg/functions"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer("okctl-mcp-server", "0.0.1")

	toolSets := functions.InitToolsets()
	toolSets.RegisterTools(s)
	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
