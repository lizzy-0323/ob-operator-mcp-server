package functions

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func listAllClusters() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool(
			"list_all_clusters",
			mcp.WithDescription("List all oceanbase clusters"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			input := "okctl cluster list"
			output, err := exec.Command("sh", "-c", input).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
			if outputStr == "" {
				return mcp.NewToolResultText("No clusters found"), nil
			}
			return mcp.NewToolResultText(outputStr), nil
		}
}

func deleteCluster() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool(
			"delete_cluster",
			mcp.WithDescription("Delete an oceanbase cluster in the specified namespace"),
			mcp.WithString("namespace",
				mcp.Description("The namespace to delete the cluster from"),
			),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("The name of the cluster to delete"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			namespace, ok := req.Params.Arguments["namespace"].(string)
			if !ok {
				namespace = "default"
			}

			clusterName, ok := req.Params.Arguments["cluster_name"].(string)
			if !ok {
				return nil, fmt.Errorf("cluster_name parameter is required")
			}

			input := fmt.Sprintf("okctl cluster delete %s -n %s", clusterName, namespace)
			output, err := exec.Command("sh", "-c", input).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
			return mcp.NewToolResultText(outputStr), nil
		}
}

func createCluster() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool(
			"create_cluster",
			mcp.WithDescription("Create a new oceanbase cluster in the specified namespace"),
			mcp.WithString("namespace",
				mcp.Required(),
				mcp.Description("The namespace to create the cluster in"),
			),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("The name of the cluster to create"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			namespace, ok := req.Params.Arguments["namespace"].(string)
			if !ok {
				namespace = "default"
			}

			clusterName, ok := req.Params.Arguments["cluster_name"].(string)
			if !ok {
				return nil, fmt.Errorf("cluster_name parameter is required")
			}

			input := fmt.Sprintf("okctl cluster create %s -n %s", clusterName, namespace)
			output, err := exec.Command("sh", "-c", input).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
			return mcp.NewToolResultText(outputStr), nil
		}
}
