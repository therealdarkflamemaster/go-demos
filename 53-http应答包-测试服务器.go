package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	// w : 写回给客户端（浏览器）， r : 从客户端（浏览器）读到的数据
	w.Write([]byte("hello world"))
}

func main() {

	// 注册回调函数，该回调函数会在服务器被访问是自动被调用
	http.HandleFunc("/itcast", handler)

	// 绑定服务器的监听地址
	http.ListenAndServe("127.0.0.1:8000", nil)

}
