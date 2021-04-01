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

func WriteMsgToClient(clnt Client, conn net.Conn) {
	// 监听 用户自带channel上是否有消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func HandlerConnect1(conn net.Conn) {
	defer conn.Close()
	// 获取用户的网络地址，ip+port
	netAddr := conn.RemoteAddr().String()
	// 创建新用户的结构体信息
	clnt := Client{make(chan string), netAddr, netAddr}

	// 将新连接的用户添加到在线用户map, key = ip+port, value = client
	onlineMap[netAddr] = clnt

	// 创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(clnt, conn)

	// 发送用户上线消息到message通道
	message <- "[" + netAddr + "]" + clnt.Name + " login ..."

	// 保证不退出
	for {

	}

}

func Manager() {
	// 初始化 onlineMap
	onlineMap := make(map[string]Client)

	for {
		// 循环全局channel中是否有数据, 有数据存在msg，无数据堵塞
		msg := <-message

		// 循环发送消息给所有在线用户
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}

}

func main() {

	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err :", err)
		return
	}
	defer listener.Close()

	// 创建go程 manager， 管理map和全局channel
	go Manager()

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