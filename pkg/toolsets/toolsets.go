package toolsets

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ToolSet struct {
	Name        string
	Description string
	Enabled     bool
	ReadOnly    bool
	writeTools  []server.ServerTool
	readTools   []server.ServerTool
}

type ToolSetsGroup struct {
	ToolSets     map[string]*ToolSet
	readOnly     bool
	everythingOn bool
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
		Enabled:     true,
		ReadOnly:    false,
	}
}

func (ts *ToolSet) AddWriteTools(tools ...server.ServerTool) *ToolSet {
	if !ts.ReadOnly {
		ts.writeTools = append(ts.writeTools, tools...)
	}
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

func (tsg *ToolSetsGroup) IsEnabled(name string) bool {
	if tsg.everythingOn {
		return true
	}

	ts, exists := tsg.ToolSets[name]
	if !exists {
		return false
	}
	return ts.Enabled
}

func (tsg *ToolSetsGroup) EnableToolSets(names []string) error {
	// Special case for "all"
	for _, name := range names {
		if name == "all" {
			tsg.everythingOn = true
			break
		}
		err := tsg.EnableToolset(name)
		if err != nil {
			return err
		}
	}
	// Do this after to ensure all toolsets are enabled if "all" is present anywhere in list
	if tsg.everythingOn {
		for name := range tsg.ToolSets {
			err := tsg.EnableToolset(name)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func (tsg *ToolSetsGroup) EnableToolset(name string) error {
	ts, exists := tsg.ToolSets[name]
	if !exists {
		return fmt.Errorf("toolset %s does not exist", name)
	}
	ts.Enabled = true
	tsg.ToolSets[name] = ts
	return nil
}

func (ts *ToolSet) RegisterTools(s *server.MCPServer) {
	if !ts.Enabled {
		return
	}
	for _, tools := range ts.readTools {
		s.AddTool(tools.Tool, tools.Handler)
	}

	if !ts.ReadOnly {
		for _, tools := range ts.writeTools {
			s.AddTool(tools.Tool, tools.Handler)
		}
	}
}
