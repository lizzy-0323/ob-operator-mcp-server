# 导入okctl模块的mcp包
from typing import Literal
from okctl import mcp


def main(transport: Literal["stdio", "sse"] = "stdio"):
    """Main entry point of the MCP server."""
    mcp.run(transport=transport)


if __name__ == "__main__":
    main()
