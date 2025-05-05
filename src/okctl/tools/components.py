import subprocess
from typing import Optional
from okctl.utils.errors import format_error

# 导入mcp实例
from okctl import mcp

# 组件安装和更新相关的工具


@mcp.tool()
def check_component(component_name: str):
    """检查目前是否安装了ob-operator和okctl，一般来说，当第一次使用okctl-mcp-server的相关工具时，首先可以检查是否安装了ob-operator和okctl

    Args:
        component_name: 组件名称，目前支持ob-operator和okctl
    """
    try:
        if not component_name:
            return "必须指定组件名称"
        if component_name not in [
            "ob-operator",
            "okctl",
        ]:
            return f"不支持检查{component_name}组件"
        if component_name == "ob-operator":
            cmd = "kubectl get pods -n oceanbase-system"
            result = subprocess.run(
                ["sh", "-c", cmd], capture_output=True, text=True, check=True
            )
            return result.stdout
        if component_name == "okctl":
            cmd = "okctl version"
            result = subprocess.run(
                ["sh", "-c", cmd], capture_output=True, text=True, check=True
            )
            return result.stdout
    except subprocess.CalledProcessError as e:
        return format_error(e)


@mcp.tool()
def install_component(
    component_name: Optional[str] = None,
    version: Optional[str] = None,
):
    """安装OceanBase组件, 目前支持ob-operator，ob-dashboard, local-path-provisioner,cert-manager,不支持其他组件，
    如果未指定，默认将安装ob-operator和 ob-dashboard

    Args:
        component_name: 组件名称
        version: 组件版本
    """
    if component_name and component_name not in [
        "ob-operator",
        "ob-dashboard",
        "local-path-provisioner",
        "cert-manager",
    ]:
        return f"不支持安装{component_name}组件"
    try:
        cmd = f"okctl install {component_name}"

        # 添加可选参数
        if version:
            cmd += f" --version {version}"

        result = subprocess.run(
            ["sh", "-c", cmd], capture_output=True, text=True, check=True
        )
        return result.stdout
    except subprocess.CalledProcessError as e:
        return format_error(e)


@mcp.tool()
def update_component(
    component_name: Optional[str] = None,
):
    """更新OceanBase组件, 目前支持ob-operator，ob-dashboard, local-path-provisioner,cert-manager,不支持其他组件，
    如果未指定，默认将更新ob-operator和 ob-dashboard

    Args:
        component_name: 组件名称
    """
    if component_name and component_name not in [
        "ob-operator",
        "ob-dashboard",
        "local-path-provisioner",
        "cert-manager",
    ]:
        return f"不支持更新{component_name}组件"
    try:
        cmd = f"okctl update {component_name}"

        result = subprocess.run(
            ["sh", "-c", cmd], capture_output=True, text=True, check=True
        )
        return result.stdout
    except subprocess.CalledProcessError as e:
        return format_error(e)


@mcp.tool()
def install_okctl():
    """安装okctl"""
    try:
        cmd = "curl -sL https://raw.githubusercontent.com/oceanbase/ob-operator/master/scripts/install-okctl.sh | bash && ./okctl install && mv ./okctl /usr/local/bin"
        result = subprocess.run(
            ["sh", "-c", cmd], capture_output=True, text=True, check=True
        )
        return result.stdout
    except subprocess.CalledProcessError as e:
        return format_error(e)
