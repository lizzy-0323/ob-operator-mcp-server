import argparse
import logging

from okctl import mcp

# These imports are required to register tools with MCP server
# pylint: disable=unused-import
from okctl.tools import (
    backup_policy,
    clusters,
    components,
    tenants,
)

# pylint: enable=unused-import

# Configure logging
logging.basicConfig(level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s")
logger = logging.getLogger("okctl_mcp_server")


def main() -> None:
    """Main entry point for the MCP server."""
    parser = argparse.ArgumentParser(description="OceanBase cluster management tool MCP server")
    parser.add_argument("--use-sse", action="store_true", help="Use Server-Sent Events (SSE) transport")
    parser.add_argument("--port", type=int, default=8000, help="Port for SSE transport (default: 8000)")
    args = parser.parse_args()

    if args.use_sse:
        logger.info("Starting server with SSE on port %s", args.port)
        mcp.run(transport="sse", port=args.port)
    else:
        logger.info("Starting server with stdio transport")
        mcp.run(transport="stdio")


if __name__ == "__main__":
    main()
