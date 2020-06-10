package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i // 通过指针 p 读取 i
	fmt.Println(*p)
	*p = 21 // 通过指针 p 设置 i
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
