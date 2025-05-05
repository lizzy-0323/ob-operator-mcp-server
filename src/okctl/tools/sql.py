import subprocess
import os
import logging
from typing import Optional, Dict, Any
from mysql.connector import connect, Error
from okctl.utils import format_error

# 导入mcp实例
from okctl import mcp

# 配置日志
logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)
logger = logging.getLogger("oceanbase_mcp_server")

# 全局配置
global_config = None


@mcp.tool()
def configure_cluster_connection(
    cluster_name: str,
    database: str,
    namespace: str = "default",
    user: str = None,
    password: str = None,
    port: int = 2881,
) -> Dict[str, Any]:
    """
    配置集群查询相关的连接

    Args:
        cluster_name: 集群名称
        database: 数据库名称
        namespace: 命名空间，默认为default
        user: 数据库用户名，默认为root
        password: 数据库密码，
        port: 数据库端口，默认为2881

    Returns:
        数据库连接配置信息
    """
    if not cluster_name:
        raise ValueError("必须指定集群名称")
    if not database:
        raise ValueError("必须指定数据库名称")

    try:
        # 首先确认集群存在
        cmd_check = f"okctl cluster show {cluster_name} -n {namespace}"
        result_check = subprocess.run(
            ["sh", "-c", cmd_check], capture_output=True, text=True
        )
        if "not found" in result_check.stderr:
            raise ValueError(f"集群 {cluster_name} 不存在")

        # 从集群信息中提取 zone 名称
        zones = []
        for line in result_check.stdout.split("\n"):
            if line.strip() and "zone" in line.lower() and "running" in line.lower():
                parts = line.split()
                if len(parts) >= 2:
                    zones.append(parts[0].strip())

        if not zones:
            raise ValueError(f"未找到集群 {cluster_name} 的可用区信息")

        # 使用 kubectl 命令获取所有 pod 信息
        cmd = f"kubectl get pods -o wide"
        result = subprocess.run(
            ["sh", "-c", cmd], capture_output=True, text=True, check=True
        )

        if not result.stdout.strip():
            raise ValueError(f"未找到任何 POD 信息")

        # 解析 POD 信息，仅筛选与集群 zone 相关的 pod
        pod_info = result.stdout.strip().split("\n")
        pod_data = []

        for line in pod_info[1:]:  # 跳过标题行
            parts = line.split()
            if len(parts) >= 6:  # 确保有足够的列
                pod_name = parts[0]
                # 检查 pod 名称是否与任何 zone 匹配
                if any(zone in pod_name for zone in zones):
                    pod_ip = parts[5]
                    pod_data.append({"pod_name": pod_name, "pod_ip": pod_ip})

        if not pod_data:
            raise ValueError(f"未找到与集群 {cluster_name} 的 zone 相关的 POD 信息")

        # 获取第一个 pod 的 IP 地址
        ip_address = pod_data[0]["pod_ip"]
        logger.info(
            f"获取到集群IP地址: {ip_address}，来自POD: {pod_data[0]['pod_name']}"
        )

        # 配置数据库连接, 如果没有提供用户和密码, 则从环境变量中获取
        global global_config
        global_config = {
            "host": ip_address,
            "port": port,
            "user": user or os.getenv("OB_CLUSTER_USER"),
            "password": password or os.getenv("OB_CLUSTER_PASSWORD"),
            "database": database,
        }

        logger.info(
            f"数据库连接配置成功: host={ip_address}, port={port}, user={user}, database={database}"
        )

        return global_config
    except subprocess.CalledProcessError as e:
        error_msg = format_error(e)
        logger.error(f"获取集群IP地址失败: {error_msg}")
        raise ValueError(f"获取集群IP地址失败: {error_msg}")
    except Exception as e:
        logger.error(f"配置数据库连接时发生错误: {str(e)}")
        raise ValueError(f"配置数据库连接时发生错误: {str(e)}")


@mcp.tool()
def execute_cluster_sql(
    query: str,
    cluster_name: str = None,
    database: str = None,
    namespace: str = "default",
) -> str:
    """
    执行集群相关的SQL命令

    Args:
        query: SQL查询语句
        cluster_name: 集群名称，如果提供则会重新配置连接
        database: 数据库名称，如果提供则会重新配置连接
        namespace: 命名空间，默认为default

    Returns:
        查询结果
    """
    global global_config

    # 如果提供了集群名称和数据库名称，则重新配置连接
    if cluster_name and database:
        try:
            configure_cluster_connection(cluster_name, database, namespace)
        except ValueError as e:
            return f"配置数据库连接失败: {str(e)}"

    # 检查是否已配置连接
    if not global_config:
        return "未配置数据库连接，请先调用 configure_cluster_connection"

    logger.info(f"执行SQL查询: {query}")

    try:
        with connect(**global_config) as conn:
            with conn.cursor() as cursor:
                cursor.execute(query)

                # 特殊处理SHOW TABLES
                if query.strip().upper().startswith("SHOW TABLES"):
                    tables = cursor.fetchall()
                    result = [f"{global_config['database']}数据库中的表: "]  # 标题
                    result.extend([table[0] for table in tables])
                    return "\n".join(result)

                elif query.strip().upper().startswith("SHOW COLUMNS"):
                    resp_header = "表的列信息: \n"
                    columns = [desc[0] for desc in cursor.description]
                    rows = cursor.fetchall()
                    result = [",".join(map(str, row)) for row in rows]
                    return resp_header + ("\n".join([",".join(columns)] + result))

                elif query.strip().upper().startswith("DESCRIBE"):
                    resp_header = "表的描述: \n"
                    columns = [desc[0] for desc in cursor.description]
                    rows = cursor.fetchall()
                    result = [",".join(map(str, row)) for row in rows]
                    return resp_header + ("\n".join([",".join(columns)] + result))

                # 常规SELECT查询
                elif query.strip().upper().startswith("SELECT"):
                    columns = [desc[0] for desc in cursor.description]
                    rows = cursor.fetchall()
                    result = [",".join(map(str, row)) for row in rows]
                    return "\n".join([",".join(columns)] + result)

                # 非SELECT查询
                else:
                    conn.commit()
                    return f"查询执行成功。影响的行数: {cursor.rowcount}"

    except Error as e:
        logger.error(f"执行SQL '{query}'时出错: {e}")
        return f"执行查询时出错: {str(e)}"
    except Exception as e:
        logger.error(f"执行查询时发生未知错误: {str(e)}")
        return f"执行查询时发生未知错误: {str(e)}"
