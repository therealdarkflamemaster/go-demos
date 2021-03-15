package main

import (
	"fmt"
	"time"
)

//func main() {
//
//	fmt.Println("当前实际时间 :", time.Now())
//	//创建定时器
//	myTimer := time.NewTimer(time.Second*2)
//
//	nowTime := <- myTimer.C //进行读操作
//	// Timer.C 是 chan类型
//	fmt.Println("现下时间是 :",nowTime)
//}

// 3种定时的方式

//func main () {
//
//	// 1. sleep
//	time.Sleep(time.Second)
//	// 2. Timer.C
//	myTimer := time.NewTimer(time.Second*2) // 创建定时器，指定定时时长
//	nowTime := <- myTimer.C // 会自动写入系统时间
//	fmt.Println("现下时间是 :",nowTime)
//
//	// 3. time.After
//	nowTime2 := <- time.After(time.Second*2) // After返回chan
//	fmt.Println("现下时间是 :",nowTime2)
//}


// 定时器的停止和重置

func main() {
	myTimer := time.NewTimer(time.Second*3) // 创建定时器
	myTimer.Reset(1*time.Second) // 重置定时时长为1s
	go func() {
		<- myTimer.C
		fmt.Println("子go程，定时完毕")
	}()

	// myTimer.Stop() // 设置定时器停止



	for {
		;
	}


}