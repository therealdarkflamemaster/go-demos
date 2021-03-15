package main

import (
	"fmt"
	"time"
)

func main() {

	quit := make(chan bool) // 创建一个判断是否终止的channel

	fmt.Println("now  = ", time.Now())

	myTicker := time.NewTicker(time.Second * 1)

	i := 0
	go func() {
		for {
			nowTime := <-myTicker.C
			i++
			fmt.Println("now time = ", nowTime)
			if i == 3 {
				quit <- true // 解除主go程堵塞
				break // return // runtime.Goexit()
			}
		}

	}()

	<- quit // 子go程 循环获取 myTicker.C 期间， 一直堵塞


}
