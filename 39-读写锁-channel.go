package main

import (
	"fmt"
	"math/rand"
	"time"
)


func readGo39(in <-chan int, idx int) {
	for {

		num := <-in //从channel中读数据
		fmt.Printf("%dth read :%d\n", idx, num)
		time.Sleep(time.Millisecond) // 放大实验现象

	}

}

func writeGo39(out chan<- int, idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)

		out <- num // 写入channel
		fmt.Printf("%dth write : %d\n", idx, num)
		time.Sleep(time.Millisecond) // 放大实验现象

	}


}



func main () {

	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	// quit := make(chan bool) // 用于 关闭主go程的channel

	ch := make(chan int)

	for i:= 0;i<5;i++ {
		go readGo39(ch,i)
		// time.Sleep(time.Second*1)
	}

	for i:= 0;i<5;i++ {
		go writeGo39(ch,i)
		// time.Sleep(time.Second*1)
	}

	// <- quit // 不让主go程结束

	for {
		;
	}

}


