import argparse
import logging
import importlib
from typing import List

from okctl import mcp

# Configure logging
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger("okctl_mcp_server")


def load_tools(tool_names: List[str]) -> None:
    """动态加载指定的工具模块。

    Args:
        tool_names: 要加载的工具模块名称列表
    """
    for tool_name in tool_names:
        try:
            logger.info("加载工具模块: %s", tool_name)
            importlib.import_module(f"okctl.tools.{tool_name}")
        except ImportError as e:
            logger.warning("无法加载工具模块 %s: %s", tool_name, e)


def main() -> None:
    """Main entry point for the MCP server."""
    parser = argparse.ArgumentParser(
        description="OceanBase cluster management tool MCP server"
    )
    parser.add_argument(
        "--use-sse", action="store_true", help="Use Server-Sent Events (SSE) transport"
    )
    parser.add_argument(
        "--port", type=int, default=8000, help="Port for SSE transport (default: 8000)"
    )
    parser.add_argument(
        "--tools",
        type=str,
        default="all",
        help="指定要启用的工具，用逗号分隔。选项: all, clusters, tenants, backup_policy, components, sql",
    )
    args = parser.parse_args()

    # install工具默认加载
    load_tools(["install"])

    # 根据参数加载相应的工具模块
    if args.tools.lower() == "all":
        # 加载所有工具模块
        logger.info("启用所有工具")
        load_tools(["clusters", "tenants", "backup_policy", "components", "sql"])
    else:
        # 解析工具参数
        tool_modules = [module.strip().lower() for module in args.tools.split(",")]
        load_tools(tool_modules)
    if args.use_sse:
        logger.info("Starting server with SSE on port %s", args.port)
        mcp.run(transport="sse", port=args.port)
    else:
        logger.info("Starting server with stdio transport")
        mcp.run(transport="stdio")


if __name__ == "__main__":
    main()
