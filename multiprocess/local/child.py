import os
import socket
import subprocess
import sys
import time

def main():
    # 连接到本地套接字
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    sock.connect("/tmp/demo.sock")

    # 向 Go 服务器发送消息
    msg = "Hello from Python subprocess"
    sock.sendall(msg.encode("utf-8"))

    # 关闭连接
    sock.close()

if __name__ == "__main__":
    main()