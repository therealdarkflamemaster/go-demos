package main

import (
	"fmt"
	"net/http"
	"os"
)

func openSendFile(fName string, w http.ResponseWriter) {
	pathFileName := "C:/Users/shengxiangLI/Pictures" + fName
	f, err := os.Open(pathFileName)
	if err != nil {
		fmt.Println("os.Open err :", err)
		w.Write([]byte("file not exist"))
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf) // 从本地读取文件数据
		if n == 0 {
			return
		}
		w.Write(buf[:n]) // 写到浏览器上
	}

}

func myHandle2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("客户端请求 :", r.URL)
	openSendFile(r.URL.String(), w)
}

func main() {

	// 注册回调函数
	http.HandleFunc("/", myHandle2)
	// 绑定监听地址
	http.ListenAndServe("127.0.0.1:9000", nil)

}
