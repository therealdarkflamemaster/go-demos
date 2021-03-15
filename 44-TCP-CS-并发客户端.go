package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// 主动发起连接请求
	conn,err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil{
		fmt.Println("net.Dial err :", err)
		return
	}
	defer conn.Close()
	//获取键盘输入 (stdin)，将输入数据发给服务器
	go func() {
		// 不断地写入
		str := make([]byte, 4096)
		for {

			n, err := os.Stdin.Read(str)
			if err != nil{
				fmt.Println("os.Stdin.Read err :", err)
				continue
			}
			// 写给服务器, 读多少写多少
			conn.Write(str[:n])

		}
	}()
	// 回显服务器回发的大写版本
	buf := make([]byte, 4096)
	for {
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read err :", err)
			continue
		}
		fmt.Println("读到服务器数据 : ", string(buf[:n]))

	}
}
