package main

import (
	"fmt"
	"net"
)

// 创建用户结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

// 创建全局map，存储在线用户
var onlineMap map[string]Client

// 创建全局channel传递用户消息
var message = make(chan string)

func HandlerConnect1(conn net.Conn) {
	defer conn.Close()

}

func main() {

	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err :", err)
		return
	}
	defer listener.Close()

	//循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept :", err)
			return
		}

		// 启动go程处理客户端数据请求
		go HandlerConnect1(conn)
	}

}
