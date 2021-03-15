package main

import "fmt"

func main () {
	var m1 map[int] string // 定义了一个map声明，但是没有空间，不能直接存储 key-value
	// m1[1] = "Luffy"

	if m1 == nil {
		fmt.Println("m1 egale nil")
	}

	m2 := map[int]string{}
	fmt.Println(len(m2)) // 0
	fmt.Println("m2 = ", m2) // m2 =  map[]

	m2[1] = "Luffy"
	fmt.Println("m2 = ", m2) // m2 =  map[1:Luffy]

	m3 := make(map[int]string)
	fmt.Println(len(m3)) // 0
	fmt.Println("m3 = ", m3) // m3 =  map[]

	m3[1] = "Luffy"
	fmt.Println("m3 = ", m3) // m3 =  map[1:Luffy]

	// 加个容量(len)
	m4 := make(map[int]string, 5) // 不绝对，可以自动扩容
	fmt.Println(len(m4)) // 0
	fmt.Println("m4 = ", m4) // m4 =  map[]

	m4[1] = "Luffy"
	fmt.Println("m4 = ", m4) // m4 =  map[1:Luffy]


	// 初始化map

	var m5 = map[int]string {1:"Luffy", 2:"Zoro"}

	fmt.Println("m5 = ", m5)

	m6 := map[int]string {1:"Luffy", 2:"Zoro"}

	fmt.Println("m6 = ", m6)

	// 赋值

	m7 := make(map[int]string, 10)
	m7[100] = "Big Mom"
	m7[101] = "KaiDuo"
	fmt.Println("m7 = ", m7)


	// 遍历map

	for k,v := range m5 {
		fmt.Printf("key:%d --- value:%q\n",k,v)
	}

	for k := range m5 { // 默认可以省略value打印，但是不可以省略key
		fmt.Printf("key:%d",k)
	}


	// 判断 map 中的 key 是否存在
	if v,has := m5[100]; has { // m5[key] 返回 1.value 2. boolean代表是否存在
		fmt.Println("value = ", v, " has = ", has)
	}else {
		fmt.Println("value = ", v, " has = ", has)
		// value =    has =  false
	}

}


