package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1 // 将封装函数内部的错误传出，给调用者
	}
	defer resp.Body.Close()
	// 循环读取网页数据，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		// 累加每一次循环读到的buf数据，存入 result，然后一次性返回
		result += string(buf[:n])

	}
	return

}

// 爬取页面操作
func working(url string) {
	fmt.Println("正在爬取热榜页 ... ")

	// 爬取此页数据
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err :", err)
		return
	}
	fmt.Println("result = ", result)

	// 将读到的网页数据，保存成一个文件
	f, err := os.Create("v2ex 热榜抓取实验.html")
	if err != nil {
		fmt.Println("os.Create err :", err)
		return
	}
	f.WriteString(result)
	// 保存好一个文件，关闭一个文件
	defer f.Close() // 整个函数结束才会执行

}

func main() {
	// 指定爬取起始页
	var url string
	fmt.Println("请输入v2ex热榜连接 :")
	fmt.Scan(&url)

	working(url)
}
