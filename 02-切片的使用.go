package main

import "fmt"

//func main() {
//	arr := [10]int {1,2,3,4,5,6,7,8,9,10}
//
//	//s := arr[1:3:5] // 两个元素，5-1个位置
//	//
//	//fmt.Println("s = ",s)
//	//fmt.Println("s.len = ",len(s))
//	//fmt.Println("s.cap = ", cap(s))
//
//
//	s := arr[1:5:7]
//	// 如果未指定切片容量，容量跟随原数组
//	fmt.Println("s = ",s)
//	fmt.Println("s.len = ",len(s)) // 5-1 =4
//	fmt.Println("s.cap = ", cap(s)) // 7-1 = 6
//
//
//	s2 := s[0:6]
//	fmt.Println("s2 = ",s2)
//	fmt.Println("s2.len = ",len(s2)) // 6-0 = 6
//	fmt.Println("s2.cap = ", cap(s2)) // 6-0 = 6
//	// 6 是跟随原数组（s）的容量大小
//
//
//}

//func main() {
//	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
//	s := arr[2:5] // {3,4,5} 范围是左闭右开的
//	fmt.Println("s = ",s)
//	// s容量和数组arr相同
//
//	s2 := s[2:7] // {5,6,7,8,9}
//	fmt.Println("s2 = ",s2)
//	fmt.Println("s2.len = ",len(s2))
//	fmt.Println("s2.cap = ", cap(s2)) // 8-2 = 6
//}

func main() {
	//创建slice
	//1. 自动推导赋值
	s1 := []int {1,2,4,6}
	fmt.Println("s1 = ", s1)


	//s2 := make([]int, 5, 10)
	//fmt.Println("len s2 = ", len(s2), "cap = ", cap(s2))
	//
	//s3 := make([]int, 7)
	//fmt.Println("len s3 = ", len(s3), "cap = ", cap(s3))


	s1 = append(s1, 999)
	s1 = append(s1, 999)
	s1 = append(s1, 999)
	s1 = append(s1, 999)
	fmt.Println("s1 apres append= ", s1)






}