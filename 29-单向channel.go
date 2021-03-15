package main

import "fmt"

//func main()  {
//
//	ch := make(chan int) // 双向channel
//
//	var sendCh chan <- int = ch
//
//	sendCh <- 789
//
//	var recvCh <- chan int = ch
//	num := <- recvCh
//	fmt.Println("num = ", num)
//
//
//	// 反向赋值
//	// var ch2 chan int = sendCh
//}


func send(out chan <- int) {
	out <- 89
	close(out)
}

func recv(in <- chan  int) {
	n := <- in
	fmt.Println("n = ", n)
}


func main() {
	ch := make(chan int)

	go func() {
		send(ch) // 双向channel转为 写channel

	}()

	recv(ch)
}