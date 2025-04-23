# 导入okctl模块的mcp包
from okctl import mcp

# 导入所有工具模块，这会自动注册工具
import okctl.clusters
import okctl.tenants
import okctl.backup_policy
import okctl.components

def main():
    print("okctl-mcp-server-py 已启动")
    # 启动FastMCP服务器
    mcp.run()

if __name__ == "__main__":
    main()
