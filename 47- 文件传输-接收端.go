package main

import (
	"fmt"
	"net"
	"os"
)

func recvFile(conn net.Conn, fileName string) {
	// 按照文件名，创建新的文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err :", err)
		return
	}
	defer f.Close()

	// 从网络中读数据写入本地文件
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("接受文件完毕")
			return
		}
		// 写入本地文件，读多少写多少
		f.Write(buf[:n])

	}

}

func main() {
	// 创建用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Listen err :", err)
		return
	}
	defer listener.Close()

	//阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err :", err)
		return
	}
	defer conn.Close()

	// 获取文件名
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err :", err)
		return
	}
	fileName := string(buf[:n])
	// 回写ok给发送端
	conn.Write([]byte("ok"))

	// 获取文件内容
	recvFile(conn, fileName)

}
