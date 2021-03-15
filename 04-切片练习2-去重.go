package main

import "fmt"


func eliminateDuplicate(data []string ) []string {
	out := data[0:1] // 截取第一个元素

	for _,word := range data {
		ifis := true
		for _,word2 := range out {
			if word2 == word {
				ifis = false
				break
			}
		}
		if ifis {
			out = append(out, word)
		}
	}


	return out

}


func main() {

	data := []string {"red","black", "red", "pink","blue","pink","blue"}
	afterData := eliminateDuplicate(data)
	fmt.Println("After change = ", afterData)
}
