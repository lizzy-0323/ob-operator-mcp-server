package toolsets

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ToolSet struct {
	Name        string
	Description string
	Enable      bool
	ReadOnly    bool
	writeTools  []server.ServerTool
	readTools   []server.ServerTool
}

type ToolSetsGroup struct {
	ToolSets map[string]*ToolSet
	readOnly bool
}

func NewServerTool(tool mcp.Tool, handler server.ToolHandlerFunc) server.ServerTool {
	return server.ServerTool{
		Tool:    tool,
		Handler: handler,
	}
}
func NewToolSetGroup() *ToolSetsGroup {
	return &ToolSetsGroup{
		ToolSets: make(map[string]*ToolSet),
		readOnly: false,
	}
}

func NewToolSet(name string, description string) *ToolSet {
	return &ToolSet{
		Name:        name,
		Description: description,
		Enable:      false,
	}
}

func (ts *ToolSet) AddWriteTools(tools ...server.ServerTool) *ToolSet {
	ts.writeTools = append(ts.writeTools, tools...)
	return ts
}

func (ts *ToolSet) setReadOnly() {
	ts.ReadOnly = true
}

func (tsg *ToolSetsGroup) AddToolSet(ts *ToolSet) {
	if tsg.readOnly {
		ts.setReadOnly()
	}
	tsg.ToolSets[ts.Name] = ts
}

func (ts *ToolSet) AddReadTools(tools ...server.ServerTool) *ToolSet {
	ts.readTools = append(ts.readTools, tools...)
	return ts
}

func (tsg *ToolSetsGroup) RegisterTools(s *server.MCPServer) {
	for _, ts := range tsg.ToolSets {
		ts.RegisterTools(s)
	}
}

func (ts *ToolSet) RegisterTools(s *server.MCPServer) {
	for _, tools := range ts.readTools {
		s.AddTool(tools.Tool, tools.Handler)
	}

	if !ts.ReadOnly {
		for _, tools := range ts.writeTools {
			s.AddTool(tools.Tool, tools.Handler)
		}
	}
}
