package okctl

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

func showCluster() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool(
			"show_cluster",
			mcp.WithDescription("Show overview of an ob cluster"),
			mcp.WithString("namespace",
				mcp.Description("The namespace of the ob cluster (default \"default\")"),
			),
			mcp.WithString("cluster_name",
				mcp.Required(),
				mcp.Description("The name of the cluster to show"),
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

			cmd := fmt.Sprintf("okctl cluster show %s -n %s", clusterName, namespace)
			output, err := exec.Command("sh", "-c", cmd).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
			return mcp.NewToolResultText(outputStr), nil
		}
}

func scaleCluster() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool(
			"scale_cluster",
			mcp.WithDescription("Scale an ob cluster, support add/adjust/delete of zones"),
			mcp.WithString("namespace",
				mcp.Description("namespace of ob cluster (default \"default\")"),
			),
			mcp.WithString("cluster_name",
				mcp.Required(),
				mcp.Description("The name of the cluster to scale"),
			),
			mcp.WithString("zones",
				mcp.Required(),
				mcp.Description("The zone of the cluster, e.g. 'z1=1', set replicas to 0 to delete the zone"),
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

			zones, ok := req.Params.Arguments["zones"].(string)
			if !ok {
				return nil, fmt.Errorf("zones parameter is required")
			}

			cmd := fmt.Sprintf("okctl cluster scale %s -n %s --zones=%s", clusterName, namespace, zones)
			output, err := exec.Command("sh", "-c", cmd).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
			return mcp.NewToolResultText(outputStr), nil
		}
}

func updateCluster() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool(
			"update_cluster",
			mcp.WithDescription("Update an ob cluster, support cpu/memory/storage"),
			mcp.WithString("namespace",
				mcp.Description("namespace of ob cluster (default \"default\")"),
			),
			mcp.WithString("cluster_name",
				mcp.Required(),
				mcp.Description("The name of the cluster to update"),
			),
			mcp.WithString("cpu",
				mcp.Description("The cpu of the observer (default 2)"),
			),
			mcp.WithString("memory",
				mcp.Description("The memory of the observer (default 10)"),
			),
			mcp.WithString("data_storage_class",
				mcp.Description("The storage class of the data storage"),
			),
			mcp.WithString("data_storage_size",
				mcp.Description("The size of the data storage (default 50)"),
			),
			mcp.WithString("log_storage_class",
				mcp.Description("The storage class of the log storage"),
			),
			mcp.WithString("log_storage_size",
				mcp.Description("The size of the log storage (default 20)"),
			),
			mcp.WithString("redo_log_storage_class",
				mcp.Description("The storage class of the redo log storage"),
			),
			mcp.WithString("redo_log_storage_size",
				mcp.Description("The size of the redo log storage (default 50)"),
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

			cmd := fmt.Sprintf("okctl cluster update %s -n %s", clusterName, namespace)

			// Add optional parameters if they are provided
			if val, ok := req.Params.Arguments["cpu"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --cpu %s", val)
			}
			if val, ok := req.Params.Arguments["memory"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --memory %s", val)
			}
			if val, ok := req.Params.Arguments["data_storage_class"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --data-storage-class %s", val)
			}
			if val, ok := req.Params.Arguments["data_storage_size"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --data-storage-size %s", val)
			}
			if val, ok := req.Params.Arguments["log_storage_class"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --log-storage-class %s", val)
			}
			if val, ok := req.Params.Arguments["log_storage_size"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --log-storage-size %s", val)
			}
			if val, ok := req.Params.Arguments["redo_log_storage_class"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --redo-log-storage-class %s", val)
			}
			if val, ok := req.Params.Arguments["redo_log_storage_size"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --redo-log-storage-size %s", val)
			}

			output, err := exec.Command("sh", "-c", cmd).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
			return mcp.NewToolResultText(outputStr), nil
		}
}

func upgradeCluster() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool(
			"upgrade_cluster",
			mcp.WithDescription("Upgrade an ob cluster, please specify the new image"),
			mcp.WithString("namespace",
				mcp.Description("namespace of ob cluster (default \"default\")"),
			),
			mcp.WithString("cluster_name",
				mcp.Required(),
				mcp.Description("The name of the cluster to upgrade"),
			),
			mcp.WithString("image",
				mcp.Required(),
				mcp.Description("The image of observer"),
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

			image, ok := req.Params.Arguments["image"].(string)
			if !ok {
				return nil, fmt.Errorf("image parameter is required")
			}

			cmd := fmt.Sprintf("okctl cluster upgrade %s -n %s --image %s", clusterName, namespace, image)
			output, err := exec.Command("sh", "-c", cmd).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
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
			mcp.WithString("cluster_name",
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
				mcp.Description("The namespace to create the cluster in (default \"default\")"),
			),
			mcp.WithString("cluster_name",
				mcp.Required(),
				mcp.Description("The name of the cluster to create"),
			),
			mcp.WithString("backup_storage_address",
				mcp.Description("The storage class of the backup storage"),
			),
			mcp.WithString("backup_storage_path",
				mcp.Description("The size of the backup storage"),
			),
			mcp.WithString("cpu",
				mcp.Description("The cpu of the observer (default 2)"),
			),
			mcp.WithString("data_storage_class",
				mcp.Description("The storage class of the data storage (default \"local-path\")"),
			),
			mcp.WithString("data_storage_size",
				mcp.Description("The size of the data storage (default 50)"),
			),
			mcp.WithString("id",
				mcp.Description("The id of the cluster"),
			),
			mcp.WithString("image",
				mcp.Description("The image of the observer (default \"quay.io/oceanbase/oceanbase-cloud-native:4.3.3.1-101000012024102216\")"),
			),
			mcp.WithString("log_storage_class",
				mcp.Description("The storage class of the log storage (default \"local-path\")"),
			),
			mcp.WithString("log_storage_size",
				mcp.Description("The size of the log storage (default 20)"),
			),
			mcp.WithString("memory",
				mcp.Description("The memory of the observer (default 10)"),
			),
			mcp.WithString("mode",
				mcp.Description("The mode of the cluster (default \"service\")"),
			),
			mcp.WithString("parameters",
				mcp.Description("Other parameter settings in OBCluster, e.g., __min_full_resource_pool_memory"),
			),
			mcp.WithString("redo_log_storage_class",
				mcp.Description("The storage class of the redo log storage (default \"local-path\")"),
			),
			mcp.WithString("redo_log_storage_size",
				mcp.Description("The size of the redo log storage (default 50)"),
			),
			mcp.WithString("root_password",
				mcp.Description("The root password of the cluster"),
			),
			mcp.WithString("zones",
				mcp.Description("The zones of the cluster, e.g. '--zones=<zone>=<replica>' (default [z1=1])"),
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

			cmd := fmt.Sprintf("okctl cluster create %s -n %s", clusterName, namespace)

			// Add optional parameters if they are provided
			if val, ok := req.Params.Arguments["backup_storage_address"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --backup-storage-address %s", val)
			}
			if val, ok := req.Params.Arguments["backup_storage_path"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --backup-storage-path %s", val)
			}
			if val, ok := req.Params.Arguments["cpu"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --cpu %s", val)
			}
			if val, ok := req.Params.Arguments["data_storage_class"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --data-storage-class %s", val)
			}
			if val, ok := req.Params.Arguments["data_storage_size"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --data-storage-size %s", val)
			}
			if val, ok := req.Params.Arguments["id"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --id %s", val)
			}
			if val, ok := req.Params.Arguments["image"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --image %s", val)
			}
			if val, ok := req.Params.Arguments["log_storage_class"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --log-storage-class %s", val)
			}
			if val, ok := req.Params.Arguments["log_storage_size"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --log-storage-size %s", val)
			}
			if val, ok := req.Params.Arguments["memory"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --memory %s", val)
			}
			if val, ok := req.Params.Arguments["mode"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --mode %s", val)
			}
			if val, ok := req.Params.Arguments["parameters"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --parameters %s", val)
			}
			if val, ok := req.Params.Arguments["redo_log_storage_class"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --redo-log-storage-class %s", val)
			}
			if val, ok := req.Params.Arguments["redo_log_storage_size"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --redo-log-storage-size %s", val)
			}
			if val, ok := req.Params.Arguments["root_password"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --root-password %s", val)
			}
			if val, ok := req.Params.Arguments["zones"].(string); ok && val != "" {
				cmd += fmt.Sprintf(" --zones %s", val)
			}

			output, err := exec.Command("sh", "-c", cmd).Output()
			if err != nil {
				return nil, fmt.Errorf("failed to run command: %v", err)
			}
			outputStr := string(output)
			return mcp.NewToolResultText(outputStr), nil
		}
}
