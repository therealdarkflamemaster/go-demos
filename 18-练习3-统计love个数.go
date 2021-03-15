package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(fileName string) int {
	fp, err := os.Open(fileName)
	if err != nil {
		fmt.Println("readFile err :", err)
		return -1
	}
	defer fp.Close()

	row := bufio.NewReader(fp)
	var total int = 0
	for{
		buf,err := row.ReadBytes('\n')
		if err != nil {
			fmt.Println(" err in readBytes : ", err)
			break
		}
		total += wordCount(string(buf[:]))
	}
	return total
}

func wordCount(s string) int {
	arr := strings.Fields(s)
	m := make(map[string]int)

	for i:=0;i<len(arr);i++ {
		m[arr[i]]++
	}
	for k,v := range m {
		if k == "love" {
			fmt.Printf("%s : %d\n", k,v)
			return m[k]
		}
	}

	return 0 // 没有love
}

func main() {
	fmt.Println("输入要查找的目录: ")
	var path string
	fmt.Scan(&path)

	dir,_ := os.OpenFile(path, os.O_RDONLY, os.ModeDir)

	defer dir.Close()

	names, _ := dir.Readdir(-1)

	var AllLove int = 0
	for _,name := range names {
		if !name.IsDir() {
			s := name.Name()
			if strings.HasSuffix(s,".txt") {
				AllLove += readFile(path + s) // 记住要路径拼接
			}
		}
	}

	fmt.Printf("里面一共有 %d 个love", AllLove)
}
