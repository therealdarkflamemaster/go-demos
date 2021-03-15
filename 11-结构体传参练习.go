package main

import "fmt"

type Person3 struct {
	name string
	age int
	flg bool
	interests []string
}

// 通过函数的参数 初始化 结构体
func initFunc(p *Person3) { // 需要地址传递，因为是初始化函数
	p.name = "name22"
	p.age = 22
	p.flg = true
	p.interests = append(p.interests, "shopping")
	p.interests = append(p.interests, "sleeping")
}


// 通过函数返回值 初始化 结构体
func initFunc2() *Person3 { // 需要地址传递，因为是初始化函数
	p := new(Person3) // p是局部变量
	p.name = "name22"
	p.age = 22
	p.flg = true
	p.interests = append(p.interests, "shopping")
	p.interests = append(p.interests, "sleeping")
	return p // 放回的是 局部变量的值
}

func main () {

	var person Person3
	initFunc(&person)
	fmt.Println("person = ", person)
	fmt.Println("person2 = ", initFunc2())
}
