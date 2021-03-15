package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwmutex sync.RWMutex // 锁只有一把，两个属性

func readGo(in <-chan int, idx int) {
	for {
		rwmutex.RLock() // 以读模式加锁
		num := <-in
		fmt.Printf("%dth read :%d\n", idx, num)
		time.Sleep(time.Millisecond) // 放大实验现象
		rwmutex.RUnlock() // 以读模式解锁
	}

}

func writeGo(out chan<- int, idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwmutex.Lock() // 以写模式加锁
		out <- num
		fmt.Printf("%dth write : %d\n", idx, num)
		time.Sleep(time.Millisecond) // 放大实验现象
		rwmutex.Unlock() //以写模式解锁
	}


}



func main () {

	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	// quit := make(chan bool) // 用于 关闭主go程的channel
	ch := make(chan int) // 用于数据传递的channel

	for i:= 0;i<5;i++ {
		go readGo(ch,i)
		// time.Sleep(time.Second*1)
	}

	for i:= 0;i<5;i++ {
		go writeGo(ch,i)
		// time.Sleep(time.Second*1)
	}

	// <- quit // 不让主go程结束

	for {
		;
	}

}
