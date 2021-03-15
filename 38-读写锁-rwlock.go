package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwmutex38 sync.RWMutex // 锁只有一把，两个属性

var value int // 定义全局变量，模拟共享数据

func readGo38(idx int) {
	for {
		rwmutex38.RLock() // 以读模式加锁
		num := value
		fmt.Printf("%dth read :%d\n", idx, num)
		time.Sleep(time.Millisecond) // 放大实验现象
		rwmutex38.RUnlock() // 以读模式解锁
	}

}

func writeGo38(idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwmutex38.Lock() // 以写模式加锁
		value = num
		fmt.Printf("%dth write : %d\n", idx, num)
		time.Sleep(time.Millisecond) // 放大实验现象
		rwmutex38.Unlock() //以写模式解锁
	}


}



func main () {

	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	// quit := make(chan bool) // 用于 关闭主go程的channel


	for i:= 0;i<5;i++ {
		go readGo38(i)
		// time.Sleep(time.Second*1)
	}

	for i:= 0;i<5;i++ {
		go writeGo38(i)
		// time.Sleep(time.Second*1)
	}

	// <- quit // 不让主go程结束

	for {
		;
	}

}

