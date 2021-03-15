package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用channel完成同步
//var ch = make(chan int)
//
//func printer37(str string) {
//	for _,ch := range str {
//		fmt.Printf("%c", ch)
//		time.Sleep(time.Millisecond*300)
//	}
//
//}
//
//func person371() {  // 先执行
//	printer37("hello")
//	ch <- 98
//
//}
//
//func person372() {  // 后执行
//	<-ch
//	printer37("world")
//}
//
//
//func main() {
//	go person371()
//	go person372()
//	for {
//		;
//	}
//}

// 使用锁完成同步

var mutex sync.Mutex // 创建一个互斥量，默认状态是 0，未加锁状态. 锁只有一把
func printer37(str string) {
	mutex.Lock() // 访问共享数据前，加锁
	for _,ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond*300)
	}
	mutex.Unlock() // 访问结束，解锁
}
func person371() {  // 先执行
	printer37("hello")

}
func person372() {  // 后执行
	printer37("world")
}
func main() {
	go person371()
	go person372()
	for {
		;
	}
}
