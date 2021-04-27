package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

// 先爬取指定url的页面， 返回一个result
func HttpGET(url string) (result string, err error) {

	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	buf := make([]byte, 4096)
	// 循环爬取整页数据
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return

}

func SpiderPage(index int) {
	// 获取url
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((index-1)*25) + "&filter="

	// 封装 HttpGET 爬取url对应页面
	result, err := HttpGET(url)
	if err != nil {
		fmt.Println("HttpGet err: ", err)
		return
	}

	fmt.Println("result = ", result)

	// 解析,编译正则表达式 -- 电影名称
	ret := regexp.MustCompile(`<img width="100"  alt="(?s:(.*?))"`)
	// 提取信息
	filmNames := ret.FindAllStringSubmatch(result, -1)
	for _, name := range filmNames {
		fmt.Println("name : ", name[1])
	}

}

func toWork(start, end int) {
	fmt.Println("正在爬取 %d 到 %d 页...\n", start, end)
	for i := start; i <= end; i++ {
		SpiderPage(i)
	}
}

func main() {
	//指定爬取的起始页，终止页
	var start, end int
	fmt.Println("input the start page :")
	fmt.Scan(&start)
	fmt.Println("input the last page : ")
	fmt.Scan(&end)

	toWork(start, end)
}
