# 导入okctl模块的mcp包
from okctl import mcp

# 导入所有工具模块，这会自动注册工具
import okctl.clusters
import okctl.tenants
import okctl.backup_policy
import okctl.components
import argparse
import logging

# Configure logging
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger("oceanbase_mcp_server")


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--use_sse", action="store_true", help="use sse")
    parser.add_argument("--port", type=int, default=8000, help="port for sse")
    args = parser.parse_args()
    if args.use_sse:
        logger.info(f"Starting server with SSE on port {args.port}")
        mcp.run(transport="sse", port=args.port)
    else:
        logger.info("Starting server with stdio")
        mcp.run(transport="stdio")


if __name__ == "__main__":
    main()
