package main

import (
	"fmt"
)

func producer40(out chan<- int, index int) {
	for i:=0;i<10;i++ {
		out <- i
		fmt.Printf("producer %dth: %d\n", index,i)
	}
	close(out)

}


func consumer40(in <-chan int, index int) {
	for num := range in {
		fmt.Printf("---consumer %dth: %d\n", index,num)
		// time.Sleep(time.Millisecond*300)
	}


}

func main() {

	product := make(chan int)

	for i:=0;i<5;i++ {
		go producer40(product, i+1)  // 一个生产者
	}
	for i:=0;i<5;i++ {
		go consumer40(product, i+1)
	}
	for {
		;
	}

}