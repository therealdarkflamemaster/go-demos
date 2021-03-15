package main

import "fmt"

func mapDelete(m map[int]string, key int)map[int]string {
	delete(m, 2)
	return m // 不用返回值也可以，因为是map的引用
}


func main () {

	m := map[int] string{1:"Luffy", 2:"Zoro"}

	m2 := mapDelete(m, 2)

fmt.Println("mapDelete = ", m2)


}
