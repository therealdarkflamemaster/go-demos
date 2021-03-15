package main

import "fmt"

func noEmpty(data []string) []string {
	out := data[:0] // 在原切片上截取，长度为0
	for _, str := range data { // range有两个返回值，index和data[index]， _ 表示不使用该值
		if str != "" {
			out = append(out, str)
		}

		// 空字符串时跳过
	}

	return out

}

func noEmpty2(data []string) []string { //slice 作为 返回值和参数的时候 传递的时地址值（引用）
	i := 0 // 下标值
	for _,str := range data {
		if str != ""{
			data[i] = str
			i++
		}
		// 空字符串时跳过
	}

	return data[:i]


}

func main() {
	data := []string{"red","","black","", "", "pink", "blue"}

	// afterData := noEmpty(data)
	afterData2 := noEmpty2(data)
	// fmt.Println("afterData = ",afterData2)
	fmt.Println("afterData2 = ",afterData2)
}

