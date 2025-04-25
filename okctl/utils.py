import subprocess


def format_error(e: subprocess.CalledProcessError):
    return f"执行命令失败:\n命令: {e.cmd}\n错误信息: {e.output}"
