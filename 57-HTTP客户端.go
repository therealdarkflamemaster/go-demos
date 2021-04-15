package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 获取服务器 应答包内容
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("http.Get err :", err)
	}
	defer resp.Body.Close()

	// 简单查看 应答包
	fmt.Println("Header : ", resp.Header)
	fmt.Println("Status : ", resp.Status)
	fmt.Println("StatusCode : ", resp.StatusCode)
	fmt.Println("Proto : ", resp.Proto)

	buf := make([]byte, 4096)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("---Read finish---")
			break
		}
		if err != nil {
			fmt.Println("resp.Body.Read(buf) err : ", err)
			break
		}

		result += string(buf[:n])
	}

	fmt.Printf("|%v|\n", result)
}
