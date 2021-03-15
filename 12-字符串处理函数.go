package main

import (
	"fmt"
	"strings"
)

func main () {
	str := "I love my work and I love my family too"

	// 字符串安装指定分隔符拆分
	ret := strings.Split(str, " ") // 安装空格进行拆分
	for _,s := range ret {
		fmt.Println(s)
	}

	// 字符串安装空格拆分
	ret2 := strings.Fields(str)
	for _,s := range ret2 {
		fmt.Println(s)
	}

	// 判断字符串结束标记
	flg := strings.HasSuffix("test.abc", ".abc")
	fmt.Println(flg)
	// 判断字符串起始标记
	flg2 := strings.HasPrefix("test.abc", "test")
	fmt.Println(flg2)
}
