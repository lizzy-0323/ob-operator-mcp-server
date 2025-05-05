import subprocess
import os
import sys
from typing import Optional, List, Dict, Any, Union, Tuple
import logging
from okctl.utils.errors import format_error

# 导入mcp实例
from okctl import mcp

# 设置日志
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')
logger = logging.getLogger(__name__)

def check_command_exists(command: str) -> bool:
    """检查命令是否存在
    
    Args:
        command: 要检查的命令
        
    Returns:
        bool: 命令是否存在
    """
    try:
        result = subprocess.run(['which', command], capture_output=True, text=True)
        return result.returncode == 0
    except Exception:
        return False

def check_kubernetes_available() -> bool:
    """检查Kubernetes是否可用
    
    Returns:
        bool: Kubernetes是否可用
    """
    try:
        result = subprocess.run(['kubectl', 'version', '--client'], capture_output=True, text=True)
        return result.returncode == 0
    except Exception:
        return False

def check_component_installed(component_name: str) -> bool:
    """检查组件是否已安装
    
    Args:
        component_name: 组件名称
        
    Returns:
        bool: 组件是否已安装
    """
    if component_name == "okctl":
        return check_command_exists("okctl")
    elif component_name == "ob-operator":
        try:
            result = subprocess.run(
                ['kubectl', 'get', 'deployment', '-n', 'oceanbase', 'ob-operator'], 
                capture_output=True, 
                text=True
            )
            return result.returncode == 0
        except Exception:
            return False
    return False

# @mcp.tool()
# def install_prerequisites():
#     """安装必要的先决条件（kubectl等）"""
#     try:
#         if not check_command_exists("kubectl"):
#             logger.info("正在安装kubectl...")
#             cmd = "curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl\" && chmod +x kubectl && sudo mv kubectl /usr/local/bin/"
#             result = subprocess.run(["sh", "-c", cmd], capture_output=True, text=True, check=True)
#             logger.info("kubectl安装完成")
#             return "kubectl安装完成"
#         else:
#             return "kubectl已经安装"
#     except subprocess.CalledProcessError as e:
#         error_msg = format_error(e)
        # logger.error(f"安装kubectl失败: {error_msg}")
        # return f"安装kubectl失败: {error_msg}"

@mcp.tool()
def install_okctl():
    """安装okctl"""
    try:
        if check_component_installed("okctl"):
            logger.info("okctl已经安装")
            return "okctl已经安装"
            
        logger.info("正在安装okctl...")
        cmd = "curl -sL https://raw.githubusercontent.com/oceanbase/ob-operator/master/scripts/install-okctl.sh | bash && chmod +x ./okctl && mv ./okctl /usr/local/bin"
        result = subprocess.run(["sh", "-c", cmd], capture_output=True, text=True, check=True)
        logger.info("okctl安装完成")
        return "okctl安装完成"
    except subprocess.CalledProcessError as e:
        error_msg = format_error(e)
        logger.error(f"安装okctl失败: {error_msg}")
        return f"安装okctl失败: {error_msg}"

@mcp.tool()
def install_ob_operator():
    """安装ob-operator"""
    try:
        if check_component_installed("ob-operator"):
            logger.info("ob-operator已经安装")
            return "ob-operator已经安装"
            
        logger.info("正在安装ob-operator...")
        cmd = "kubectl apply -f https://raw.githubusercontent.com/oceanbase/ob-operator/stable/deploy/operator.yaml"
        result = subprocess.run(["sh", "-c", cmd], capture_output=True, text=True, check=True)
        logger.info("ob-operator安装完成")
        return "ob-operator安装完成"
    except subprocess.CalledProcessError as e:
        error_msg = format_error(e)
        logger.error(f"安装ob-operator失败: {error_msg}")
        return f"安装ob-operator失败: {error_msg}"

# @mcp.tool()
# def setup_environment():
#     """设置环境，安装所有必要的组件"""
#     results = []
    
#     # 安装先决条件
#     prereq_result = install_prerequisites()
#     results.append(f"先决条件安装结果: {prereq_result}")
    
#     # 安装okctl
#     okctl_result = install_okctl()
#     results.append(f"okctl安装结果: {okctl_result}")
    
#     # 安装ob-operator
#     operator_result = install_ob_operator()
#     results.append(f"ob-operator安装结果: {operator_result}")
    
#     return "\n".join(results)
