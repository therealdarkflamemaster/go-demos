package main

import (
	"fmt"
	"net"
)

func main()  {
	// 指定服务器的协议，ip地址和port端口号
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("listen function err : ", err)
		return
	}

	defer listener.Close()

	fmt.Println("等待客户端连接...")
	// 堵塞监听客户端连接请求,返回用于通信的socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept function err : ", err)
		return
	}
	defer conn.Close()

	fmt.Println("服务器与客户端建立连接成功")
	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n,err := conn.Read(buf)
	if err != nil {
		fmt.Println("Connect.Read  err : ", err)
		return
	}
	// 处理数据 -- 打印
	fmt.Println("服务器读到数据 ：", string(buf[:n]))

	// 关闭两个socket


}
