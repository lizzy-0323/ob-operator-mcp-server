import subprocess
from typing import Optional

# 导入mcp实例
from okctl import mcp


# 组件安装和更新相关的工具
@mcp.tool()
def install_component(
    component_name: str,
    version: Optional[str] = None,
):
    """安装OceanBase组件

    Args:
        component_name: 组件名称
        version: 组件版本
    """
    try:
        cmd = f"okctl install {component_name}"

        # 添加可选参数
        if version:
            cmd += f" --version {version}"

        result = subprocess.run(["sh", "-c", cmd], capture_output=True, text=True, check=True)
        return result.stdout
    except subprocess.CalledProcessError as e:
        return f"执行命令失败: {e}"


@mcp.tool()
def update_component(
    component_name: str,
):
    """更新OceanBase组件

    Args:
        component_name: 组件名称
    """
    try:
        cmd = f"okctl update {component_name}"

        result = subprocess.run(["sh", "-c", cmd], capture_output=True, text=True, check=True)
        return result.stdout
    except subprocess.CalledProcessError as e:
        return f"执行命令失败: {e}"


@mcp.tool()
def install_okctl():
    """安装okctl"""
    try:
        cmd = "curl -sL https://raw.githubusercontent.com/oceanbase/ob-operator/master/scripts/install-okctl.sh | bash && sudo mv ./okctl /usr/local/bin/okctl "
        result = subprocess.run(["sh", "-c", cmd], capture_output=True, text=True, check=True)
        return result.stdout
    except subprocess.CalledProcessError as e:
        return f"执行命令失败: {e}"
