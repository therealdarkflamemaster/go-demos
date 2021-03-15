package main

import (
	"fmt"
	"runtime"
)

func main() {


	go func() {
		for {
			fmt.Println(" this is the goroutine test")

		}

	}()

	for {
		runtime.Gosched() // 出让当前cpu时间片
		fmt.Println(" this is the main test")

	}
}