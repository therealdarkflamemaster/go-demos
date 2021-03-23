package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	// 组织一个UDP的地址结构,指定服务器的ip和port
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err :", err)
		return
	}
	fmt.Println("UDP 服务器地址结构创建完成")
	// 创建用于通信的socket
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err :", err)
		return
	}

	defer udpConn.Close()
	fmt.Println("UDP 服务器通信socket创建完成")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	// 返回三个值，分别是 读取到的字节数，客户端地址，error
	n, cltAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("udpConn.ReadFromUDP err :", err)
		return
	}
	// 模拟处理数据
	fmt.Println("来自客户端 : ", cltAddr)
	fmt.Println("服务器读到 : ", string(buf[:n]))

	// 回写数据给客户端
	// 提取系统时间
	daytime := time.Now().String()
	_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
	if err != nil {
		fmt.Println("udpConn.WriteToUDP err :", err)
		return
	}

}
