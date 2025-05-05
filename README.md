# OceanBase Kubernetes Control Tool (okctl) MCP Server

English | [简体中文](README-CN.md)

## Project Overview

This project is the MCP server implementation for the OceanBase Kubernetes Control Tool ([okctl](https://github.com/oceanbase/ob-operator?tab=readme-ov-file#using-cli-tool-okctl)). It provides a set of tool functions for managing OceanBase clusters, tenants, and backup policies. These functions are implemented by calling the underlying okctl command-line tool and exposing these features to clients through the MCP protocol.

## Functional Modules

This project includes the following main tools:

### 1. Cluster Management (clusters.py)

Provides functionality for creating, deleting, viewing, scaling, updating, and upgrading OceanBase clusters.

- `list_all_clusters()` - List all OceanBase clusters
- `show_cluster()` - Display detailed information of a specified cluster
- `create_cluster()` - Create a new OceanBase cluster
- `delete_cluster()` - Delete a specified OceanBase cluster
- `scale_cluster()` - Scale an OceanBase cluster
- `update_cluster()` - Update OceanBase cluster configuration
- `upgrade_cluster()` - Upgrade OceanBase cluster version

### 2. Tenant Management (tenants.py)

Provides functionality for creating, deleting, viewing, scaling, updating, and managing OceanBase tenants.

- `list_tenants()` - List all tenants
- `create_tenant()` - Create a new tenant
- `delete_tenant()` - Delete a specified tenant
- `show_tenant()` - Display tenant detailed information
- `scale_tenant()` - Scale tenant resources
- `update_tenant()` - Update tenant configuration
- `upgrade_tenant()` - Upgrade tenant version
- `change_tenant_password()` - Change tenant password
- `activate_tenant()` - Activate standby tenant
- `replay_tenant_log()` - Replay tenant logs
- `switchover_tenant()` - Switch between primary and standby tenants

### 3. Backup Policy Management (backup_policy.py)

Provides functionality for creating, deleting, viewing, updating, and managing OceanBase backup policies.

- `list_backup_policies()` - List all backup policies
- `create_backup_policy()` - Create a new backup policy
- `delete_backup_policy()` - Delete a specified backup policy
- `show_backup_policy()` - Display backup policy detailed information
- `update_backup_policy()` - Update backup policy
- `pause_backup_policy()` - Pause backup policy
- `resume_backup_policy()` - Resume backup policy

### 4. SQL Operations (sql.py)

Provides functionality for configuring database connections and executing SQL queries on OceanBase clusters.

- `configure_cluster_connection()` - Configure database connection to a cluster
  - Parameters: cluster_name, database, namespace (default: "default"), user, password, port (default: 2881)
  - Returns: Database connection configuration information
  
- `execute_cluster_sql()` - Execute SQL queries on a cluster
  - Parameters: query, cluster_name (optional), database (optional), namespace (default: "default")
  - Returns: Query results
  - Supports various SQL commands including SELECT, SHOW TABLES, SHOW COLUMNS, DESCRIBE, and DML statements

### 5. Component Management (components.py)

Provides functionality for installing, updating, and managing OceanBase components.

- `list_components()` - List all installed components
- `install_component()` - Install a new component
- `update_component()` - Update component

## Development Environment Setup

### Prerequisites

- Python 3.10 or higher installed
- uv package manager installed ([uv official documentation](https://github.com/astral-sh/uv))
- OceanBase Kubernetes Control Tool (okctl) installed and configured
- Kubernetes environment configured with access permissions to OceanBase clusters

### Configuring the MCP Server

```json
{
  "mcpServers": {
    "okctl-mcp-server-py": {
      "command": "uv",
      "args": ["--directory", "/path/to/okctl-mcp-server-py", "run", "main.py"]
    }
  }
}
```
#### Command Line Arguments

- `--tools`: Specify which tools to enable, comma separated. Options:
  - `all`: Enable all tools (default)
  - `cluster`: Enable cluster management tools only
  - `tenant`: Enable tenant management tools only
  - `backup`: Enable backup policy management tools only
  - `component`: Enable component management tools only
  - `sql`: Enable SQL operation tools only
  
  Example: `--tools=cluster,tenant,sql`

- `--use_sse`: Use Server-Sent Events (SSE) transport instead of stdio
- `--port`: Specify the port for SSE transport (default: 8000)

## Important Notes

- The server needs to run in an environment with access to the Kubernetes cluster
- All functions are implemented by calling the underlying okctl command-line tool, so ensure okctl is properly installed and configured
- Most functions provide a namespace parameter with a default value of "default", which can be specified as needed
- Some operations (such as deleting clusters, deleting tenants) may be irreversible, please proceed with caution
- It is recommended to perform backups before executing important operations

## Contributing

Issues and Pull Requests are welcome to improve this project.

## License

This project is licensed under the [Apache 2.0 License](LICENSE).
