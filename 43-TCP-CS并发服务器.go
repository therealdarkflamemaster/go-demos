package main

import (
	"fmt"
	"net"
	"strings"
)


func HandlerConnect (conn net.Conn) {
	defer conn.Close()
	// 获取连接的客户端的 Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接")

	// 循环读取客户端的发送数据
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err :", err)
			return
		}
		fmt.Println("服务器读到 :", buf[:n])

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {

	// 创建监听socket
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("listen err :", err)
		return
	}
	defer listener.Close()

	//监听客户端的请求
	for {
		fmt.Println("server wait for connection ...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err :", err)
			return
		}


		// 具体完成数据通信

		go HandlerConnect(conn)
	}



}
