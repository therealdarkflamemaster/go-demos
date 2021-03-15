package main

import (
	"fmt"
	"time"
)

func sing() {
	for i:=0; i<5; i++ {
		fmt.Println("----singing, song ----")
		time.Sleep(100 * time.Millisecond)
	}


}

func dance() {
	for i:=0; i<5; i++ {
		fmt.Println("----dancing by Nicolas ZhaoSI ----")
		time.Sleep(100 * time.Millisecond)
	}


}


func main() {

	go sing()
	go dance()

	for {

	}

}
