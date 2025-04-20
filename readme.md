# okctl-mcp-server

## Overview

The okctl mcp server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction)
server that provides a set of tools for managing OceanBase clusters, and it is used by the [okctl](https://github.com/oceanbase/okctl) command-line tool.

## Use Cases

- Manage OceanBase clusters
- Manage OceanBase tenants
- Manage OceanBase backup proxies
- Install and update ob-operator related tools

## Prerequisites

### Use Docker

1. To run the server in a container, you will need to have [Docker](https://www.docker.com/) installed.
2. Once Docker is installed, you will also need to ensure Docker is running.

## Installation

### Usage with VS Code

### Usage with Claude Desktop

### Build from Source

## Tool Configuration

### Available Tools

The following sets of tools are available (all are on by default):

| Toolset         | Description                             |
| --------------- | --------------------------------------- |
| `obclusters`    | OceanBase cluster management tools      |
| `obtenants`     | OceanBase tenant management tools       |
| `obbackupproxy` | OceanBase backup proxy management tools |
| `installs`      | Install ob-operator related tools       |
| `updates`       | Update ob-operator related tools        |

## Tools

### obclusters

### obtenants

### obbackupproxy

### installs

### updates

## License

The okctl mcp server is licensed under the [Apache License 2.0](https://github.com/oceanbase/okctl-mcp-server/blob/main/LICENSE).
