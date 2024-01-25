package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

// 启动子进程时传递套接字地址
// go run socket.go server
// go run socket.go client /tmp/demo.sock
// 模拟py子进程参与通信  python3 child.py

// 启动父进程     go run socket.go server
// 启动6个子进程一起向父亲传消息   go test -bench='BenchmarkRunClient' -benchtime=20x -count=3
const socketAddr = "/tmp/demo.sock"

func main() {
	cmd := os.Args[1]

	switch cmd {
	case "server":
		runServer()
	case "client":
		// 启动子进程
		// exec.Command("python3", "child.py")
		runClient()
	default:
		fmt.Println("Invalid command:", cmd)
		os.Exit(1)
	}
}

func runServer() {
	// 删除旧的 Unix 套接字，如果存在的话
	_ = os.Remove(socketAddr)

	// 创建一个 Unix 套接字监听器
	listener, err := net.Listen("unix", socketAddr)
	if err != nil {
		fmt.Println("Error listening on socket:", err)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		// 接受来自客户端的连接
		fmt.Println("accepting coon")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 处理客户端请求
		go func() {
			handleServerConnection(conn)
		}()
	}
}

func handleServerConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		reader := bufio.NewReader(conn)
		// 阻塞式读取子进程消息
		n, err := reader.Read(buf)
		if err == io.EOF {
			fmt.Println("nil")
			time.Sleep(1 * time.Second)
			continue
		}
		// 将接收到的数据转换为字符串并打印
		msg := strings.TrimSpace(string(buf[:n]))
		fmt.Println("Received from client:", msg)
		fmt.Println("i am server")
	}
}

func runClient() {
	// 创建一个 Unix 套接字连接
	// socketAddr := os.Args[2]
	fmt.Println(socketAddr)
	conn, err := net.Dial("unix", socketAddr)
	if err != nil {
		fmt.Println("Error connecting to socket:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 向服务器发送指标消息
	for i := 0; i < 50; i++ {
		_, err = conn.Write([]byte(`{"cnt":10}`))
		if err != nil {
			fmt.Println("Error sending message:", err)
			os.Exit(1)
		}
	}
}
