package main

import "fmt"

//func main() {
//
//	data := [...] int {0,1,2,3,4,5,6,7,8,9}
//	s1 := data[8:] // {8,9}
//	s2 := data[0:5] // {0,1,2,3,4}
//
//
//	// 对应位置copy
//	copy(s2, s1)
//	fmt.Println("s2 = ", s2)
//
//
//}

func remove(data []int, idx int) []int {
	copy(data[idx:], data[idx+1:])

	return data[:len(data)-1]
}

// 练习 3
func main() {

	data := []int {5,6,7,8,9}
	// 删除7 相当于 把8，9 拷贝到 7，8，9 的位置，并且舍弃9

	afterData := remove(data, 2) // 7 的 index = 2

	fmt.Println(afterData)
}