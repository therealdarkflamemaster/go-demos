package main

import "fmt"

// 自定义结构体 （一种数据类型）
type Person2 struct {
	name string
	sex byte
	age int
}

func test2 (p *Person2) {
	p.name = "Luffy"
	p.sex = 'm'
	p.age = 17
}

func main() {

	var p1 *Person2 = &Person2{"name1", 'f', 19}
	fmt.Println("p1 = ", p1)

	var p2 *Person2
	p2 = new(Person2)
	// p2 = &Person2{"name1", 'f', 19}
	fmt.Println("p2 = ",p2)

	p3 := new(Person2)
	p3.name = "Peter"
	fmt.Printf("p3 = %p\n", p3)
	fmt.Printf("&p3.name = %p",&p3.name)


	test2(p3)
	fmt.Println("p3 = ",p3) // 保存test2 的修改，因为传递的是地址
}