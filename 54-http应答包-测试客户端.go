package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc2(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		//return // 返回当前函数调用
		//runtime.Goexit() // 结束当前go程
		os.Exit(1) // 无论什么时候调用，结束当前进程
	}
}

// 装 浏览器
func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	errFunc2(err, "net.Dial err : ")

	defer conn.Close()

	httpRequest := "GET /itcast HTTP/1.1/r/nHOST:127.0.0.1:8000/r/n"

	conn.Write([]byte(httpRequest))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))

}
