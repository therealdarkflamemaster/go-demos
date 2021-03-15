package main

import (
	"fmt"
	"strings"
)

func wordCountFunc(str string) map[string]int {
	s := strings.Fields(str) // 将字符串变成 字符串slice
	m := make(map[string]int)
	// 遍历 slice
	for _,word := range s {
		if _,has := m[word]; has {
			m[word]++
		}else {
			m[word] = 1
		}
	}
	return m
}

func main() {

	str := "I love my work and I I I I love my family too"

	mRet := wordCountFunc(str)

	fmt.Println("result = ", mRet)

}
