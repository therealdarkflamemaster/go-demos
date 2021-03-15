package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) { // 类型转换成单向写channel
	for i:=0;i<10;i++ {
		out<- i*i
		fmt.Println("producer produce num = ", i*i)
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("consumer get num = ", num)
		time.Sleep(time.Second)
	}

}

func main() {

	ch := make(chan int, 5)

	go producer(ch) // 子go程 生产者

	consumer(ch) // 主go程 消费者

}