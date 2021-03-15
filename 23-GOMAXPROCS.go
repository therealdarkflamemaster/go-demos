package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOROOT())

	n := runtime.GOMAXPROCS(1) // 只用单核
	fmt.Println("n = ",n)

	n = runtime.GOMAXPROCS(2) // 双核
	fmt.Println("n = ",2)

	for {
		go fmt.Print(0) // 子go程

		fmt.Print(1) // 主go程
	}
}