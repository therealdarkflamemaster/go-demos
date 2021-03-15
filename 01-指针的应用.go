package main

import "fmt"

func swap(a,b int) {
	a,b = b,a
	fmt.Println("(swap)a = ",a, " b = ", b)

}

func swap2(x,y *int) { // 指针参数
	*x,*y = *y,*x

}

func main() {
	//var a int = 10
	//
	//// 指针变量
	//var p *int = &a
	//
	//a = 100
	//fmt.Println("a = ",a)
	//
	//*p = 250 // 相当于改变a变量
	//fmt.Println("a = ",*p)
	//
	//a = 1000
	//fmt.Println("a = ",*p)
	//
	//p = new (int)
	//
	//fmt.Printf("%s\n", *p)
	//fmt.Printf("%q\n", *p) // 打印go语言格式的字符串


	a,b := 10,20
	swap(a,b) //传的是变量值
	fmt.Println("(main-swap1)a = ",a, " b = ", b)
	swap2(&a, &b) // 传的是地址值
	fmt.Println("(main-swap2)a = ",a, " b = ", b)
}
