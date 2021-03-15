package main

//// 死锁1
//func main() {
//	ch :=  make(chan int)
//	ch <- 789
//	num := <- ch
//	fmt.Println("num = ", num)
//}

//// 死锁2
//func main() {
//	ch :=  make(chan int)
//	num := <- ch
//	fmt.Println("num = ", num)
//	go func() {
//		ch <- 789
//
//	}()
//
//}

// 死锁3
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
				case num :=<-ch1:
					ch2<-num

			}
		}
	}()

	for {
		select {
			case num := <- ch2:
				ch1<- num

		}
	}
}