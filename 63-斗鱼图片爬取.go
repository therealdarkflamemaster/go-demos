package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// 先爬取指定url的页面， 返回一个result
func HttpGET2(url string) (result string, err error) {

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

//func SpiderPage2(index int) {
//	// 获取url
//	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((index-1)*25) + "&filter="
//
//	// 封装 HttpGET 爬取url对应页面
//	result, err := HttpGET(url)
//	if err != nil {
//		fmt.Println("HttpGet err: ", err)
//		return
//	}
//
//	fmt.Println("result = ", result)
//
//	// 解析,编译正则表达式 -- 电影名称
//	ret := regexp.MustCompile(`<img width="100"  alt="(?s:(.*?))"`)
//	// 提取信息
//	filmNames := ret.FindAllStringSubmatch(result, -1)
//	for _, name := range filmNames {
//		fmt.Println("name : ", name[1])
//	}
//
//}

func main() {

	url := "https://www.douyu.com/g_yz"

	// 爬取整个页面，将整个页面的全部信息，保存在result
	result, err := HttpGET2(url)
	if err != nil {
		fmt.Println("HttpGET2 err : ", err)
		return
	}

	ret := regexp.MustCompile(`<img src="(?s:(.*?))" class="DyImg-content is-normal"`)

	// 提取每一张图片的url
	all := ret.FindAllStringSubmatch(result, 20)

	for _, imgURL := range all {
		fmt.Println(imgURL)
	}

}
