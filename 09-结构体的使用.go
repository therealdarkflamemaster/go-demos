package main

import (
	"fmt"
	"unsafe"
)

// 自定义结构体 （一种数据类型）
type Person struct {
	name string
	sex byte
	age int
}


func test (man Person) {
	man.name = "testName"
	man.age = 100
	fmt.Println("man in test = ", man)
	fmt.Println("man in test  = ", unsafe.Sizeof(man))
}

type Student struct {
	p Person
	id int
	score int
}


func main() {

	// 顺序初始化

	var man Person = Person{
		name: "Andy",
		sex: 'm',
		age: 20,
	}
	fmt.Println("man = ", man)

	// 部分初始化

	man2 := Person{
		name:"Bob",
		age: 40,
	}
	fmt.Println("man2 = ",man2)

	// 索引成员变量
	fmt.Println("man2's name is ", man2.name)


	var man3 Person
	man3.name = "Cecile"
	man3.sex = 'f'
	man3.age = 99
	fmt.Println("man3 is ", man3)

	// 结构体比较
	fmt.Println("man3 compare to man2 ", man3 == man2) // false

	// 相同类型的结构体可以赋值
	var tmp Person
	tmp = man3
	fmt.Println("tmp = ", tmp)

	// 函数内部使用结构体传参
	var temp Person
	test(temp)
	fmt.Println("temp in main before = ", unsafe.Sizeof(temp)) // 求变量占的字节数
	fmt.Println("temp aprs test = ", temp) // 不被改动
	fmt.Println("temp in main after = ", unsafe.Sizeof(temp))

	// 结构体的地址
	fmt.Println("&tmp address is ", &man3)
	fmt.Println("&tmp's name address is ", &man3.name) // 运算符 后缀 高于 前缀

}
