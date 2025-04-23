# 从主模块导入mcp实例
# 创建FastMCP实例
from fastmcp import FastMCP

# 创建全局mcp实例供所有模块使用
mcp = FastMCP("okctl-mcp-server-py", version="0.1.0")
