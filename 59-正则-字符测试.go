package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc cat 9ca azc cba aMc"
	// 解析，编译正则表达式
	ret := regexp.MustCompile(`a[^0-9a-z]c`)

	// 提取需要的信息
	all := ret.FindAllStringSubmatch(str, -1)
	fmt.Println("All :", all)
}
