package main

import "fmt"

func main () {

	ch := make(chan string)  // 无缓冲channel
	fmt.Println("len(ch) = ", len(ch), " cap(ch) = ", cap(ch))
	// len(ch) =  0  cap(ch) =  0
	// len(ch) ： channel当中剩余未读的数据个数。 cap（ch），通道的容量

	go func() {
		for i:= 0;i<2;i++ {
			fmt.Println("i = ", i)
		}
		// 通知主go程打印完毕
		 ch <- "子go打印完毕"
	}()

	str := <- ch
	fmt.Println("str = ", str)

}
