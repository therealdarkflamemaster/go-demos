package main

import (
	"fmt"
	"net/http"
)

func myHandle(w http.ResponseWriter, r *http.Request) {
	// w: 写给客户端的数据内容
	// r: 从客户端读到的内容
	w.Write([]byte("This is a web server"))
	fmt.Println("Header : ", r.Header)
	fmt.Println("URL : ", r.URL)
	fmt.Println("Method : ", r.Method)
	fmt.Println("Host : ", r.Host)
	fmt.Println("RemoteAddr : ", r.RemoteAddr)
	fmt.Println("Body : ", r.Body)
}

func main() {
	// 注册回调函数，该函数在客户端访问服务器时，自动被调用

	http.HandleFunc("/", myHandle)

	// 绑定服务器监听地址
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Println("err :", err)
	}
}
