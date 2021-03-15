package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)  //用来进行数据通信的channel

	quit := make(chan bool)  // 用来判断是否退出的channel

	go func() {  //子go程写数据
		for i := 0;i<5;i++ {
			ch <- i
			time.Sleep(time.Second)
		}

		close(ch)
		quit <- true    // 通知 主go程退出
		runtime.Goexit() // 子go程退出
	}()


	for {   // 主go程读数据
		select {
			case num := <- ch :
				fmt.Println("read num = ", num)

			case <- quit :
				// break   // break跳出select
				return
				// runtime.Goexit() 不用于主go程退出

		}

		fmt.Println("===========")
	}



}


