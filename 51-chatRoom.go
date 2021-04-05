package main

import (
	"fmt"
	"net"
	"strings"
	"time"
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

func MakeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ": " + msg
	return
}

func HandlerConnect1(conn net.Conn) {
	defer conn.Close()

	// 创建一个channel，判断用户是否活跃
	hasData := make(chan bool)

	// 获取用户的网络地址，ip+port
	netAddr := conn.RemoteAddr().String()
	// 创建新用户的结构体信息
	clnt := Client{make(chan string), netAddr, netAddr}

	// 将新连接的用户添加到在线用户map, key = ip+port, value = client
	onlineMap[netAddr] = clnt

	// 创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(clnt, conn)

	// 发送用户上线消息到message通道
	message <- MakeMsg(clnt, "log in ...")

	// 创建一个channel，用来判断是否退出
	isQuit := make(chan bool)

	// 创建一个go程, 专门处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端:%s退出\n", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err :", err)
				return
			}
			msg := string(buf[:n-1])

			// 提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list : \n"))
				// 遍历当前map， 获取在线用户
				for _, user := range onlineMap {
					userInfo := user.Addr + " : " + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}

				// 判断用户发送了 改名 命令
			} else if len(msg) >= 8 && msg[:6] == "rename" { //rename|newName
				newName := strings.Split(msg, "|")[1] // msg[8:]
				clnt.Name = newName                   // 修改结构体成员Name
				onlineMap[netAddr] = clnt             // 更新在线用户列表
				conn.Write([]byte("rename successfully \n"))

			} else {
				// 将读到的数据广播给所有在线用户, 写入到message
				message <- MakeMsg(clnt, msg)
			}

			hasData <- true

		}
	}()

	// 保证不退出
	for {
		// 监听channel上的数据流动
		select {
		case <-isQuit:
			delete(onlineMap, clnt.Addr)           // 将用户从在线用户列表移除
			message <- MakeMsg(clnt, "logout ...") // 写入用户退出消息到全局
			return

		case <-hasData:
			// 什么都不做,目的是重置下面case的计时器

		case <-time.After(time.Second * 10):
			delete(onlineMap, clnt.Addr)           // 将用户从在线用户列表移除
			message <- MakeMsg(clnt, "logout ...") // 写入用户退出消息到全局
			return

		}
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
