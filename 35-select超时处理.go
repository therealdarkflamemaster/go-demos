package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {  //子go程获取数据
		for {
			select {
				case num := <- ch :
					fmt.Println("num = ", num)
				case <- time.After(time.Second*3) :
					quit <- true
					goto label  //跳转到函数内任意位置
					// return
			}
		}
		label:
			fmt.Println("break to label ----")
	}()

	for i:=0;i<2;i++ {
		ch <-i
		time.Sleep(time.Second*2)
	}

	<-quit // 主go程,堵塞等待子go程通知退出
	fmt.Println("finish !!")


}
