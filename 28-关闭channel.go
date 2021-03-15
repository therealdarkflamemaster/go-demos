package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)


	go func() {
		for i := 0;i<5;i++ {
			ch <- i
		}
		close(ch) // 写端，写完数据主动关闭channel
		fmt.Println("子go程结束 ")
	}()

	time.Sleep(time.Second*2)
	//for {
	//	if num, ok := <- ch; ok == true {
	//		fmt.Println("读到的 num = ", num)
	//	} else {
	//		n := <- ch
	//		fmt.Println("关闭后读到 ", n)
	//		break
	//	}
	//}
	for num := range ch {
		fmt.Println("读到的 num = ", num)
	}
}