package main

import (
	"fmt"
	"net"
)

func main() {

	// 指定 服务器的IP + port 创建 用于监听的socket
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net function err : ", err)
		return
	}
	defer conn.Close()

	// 主动写数据给服务器
	conn.Write([]byte("Are you ok?"))

	// 接受服务器回发的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err : ", err)
		return
	}
	conn.Write(buf[:n]) // 读多少，写多少
	fmt.Println("server side return : ", buf[:n])

}
