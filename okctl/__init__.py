# 从主模块导入mcp实例
# 创建FastMCP实例
from fastmcp import FastMCP

# 创建全局mcp实例供所有模块使用
mcp = FastMCP("okctl-mcp-server", version="0.1.0", loglevel="ERROR")


@mcp.prompt()
def system_prompt() -> str:
    """基础概念介绍，必须加载"""
    return """okctl是 OceanBase 集群管理工具 okctl 的命令行接口，用于管理 OceanBase 集群，租户，备份策略等资源，并且支持相关组件的安装和更新。
你可以根据用户输入合理推断 API 调用顺序，但 API 地址和传参必须参考 Prompt 严格使用，不能自己臆测参数。
"""
