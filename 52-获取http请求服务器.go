package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		//return // 返回当前函数调用
		//runtime.Goexit() // 结束当前go程
		os.Exit(1) // 无论什么时候调用，结束当前进程
	}
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, " net.Listen err :")
	defer listener.Close()

	conn, err := listener.Accept()
	errFunc(err, "listener.Accept err :")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)

	if n == 0 {
		return
	}

	errFunc(err, "conn.Read err :")

	fmt.Printf("|%s|\n", string(buf[:n]))
}
