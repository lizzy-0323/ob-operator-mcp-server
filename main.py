# 导入okctl模块的mcp包
from okctl import mcp

# 导入argparse和logging
import argparse
import logging
import importlib

# Configure logging
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger("okctl_mcp_server")


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--use_sse", action="store_true", help="use sse")
    parser.add_argument("--port", type=int, default=8000, help="port for sse")
    parser.add_argument("--tools", type=str, default="all", help="specify which tools to enable, comma separated. Options: all, cluster, tenant, backup, component, sql")
    args = parser.parse_args()
    
    # 根据参数导入相应的工具模块
    if args.tools.lower() == "all":
        # 导入所有工具模块，这会自动注册工具
        logger.info("Enabling all tools")
        import okctl.clusters
        import okctl.tenants
        import okctl.backup_policy
        import okctl.components
        import okctl.sql
    else:
        # 解析工具参数
        tool_modules = args.tools.split(",")
        for module in tool_modules:
            module = module.strip().lower()
            if module == "cluster":
                logger.info("Enabling cluster tools")
                import okctl.clusters
            elif module == "tenant":
                logger.info("Enabling tenant tools")
                import okctl.tenants
            elif module == "backup":
                logger.info("Enabling backup policy tools")
                import okctl.backup_policy
            elif module == "component":
                logger.info("Enabling component tools")
                import okctl.components
            elif module == "sql":
                logger.info("Enabling sql tools")
                import okctl.sql
            else:
                logger.warning(f"Unknown tool module: {module}")
    
    if args.use_sse:
        logger.info(f"Starting server with SSE on port {args.port}")
        mcp.run(transport="sse", port=args.port)
    else:
        logger.info("Starting server with stdio")
        mcp.run(transport="stdio")


if __name__ == "__main__":
    main()
