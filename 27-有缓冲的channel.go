package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3) // 存满3个元素之前，不堵塞

	fmt.Println("len = ", len(ch), " cap = ", cap(ch))

	go func() {
		for i:= 0;i<8;i++ {
			fmt.Println("在子go程中， i ", i ,"len = ", len(ch), " cap = ", cap(ch))
			ch <- i
		}
	}()


	time.Sleep(time.Second * 3)
	for i:= 0;i<8;i++ {
		num := <- ch
		fmt.Println("在主go程中， num ", num)

	}




}
