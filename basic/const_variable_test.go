package basic

import (
	"fmt"
	"testing"
)

type myType int

const (
	// 如果显性给出定义，那么这里就会报错，n int = 3
	n = 3
)

// 如果n是隐性给出定义，那么这里就没有问题
func TestConst(t *testing.T) {
	var m myType = 4
	fmt.Println(n + m)
}

const (
	a = 5               // 默认是int
	s = "Hello, Gopher" // 默认是string
)

// 无类型常量也有默认类型
func TestDefaultType(t *testing.T) {
	n2 := a
	var i interface{} = a

	fmt.Printf("%T\n", n2)

	fmt.Printf("%T\n", i)

	i = s
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", i)
}

const (
	d = 1
	b = 3
	c
)

// const语法提供"隐式重复前一个非空表达式"
func TestDuplicate(t *testing.T) {
	fmt.Println(d, b, c)
}

// 无类型常量，有类型常量
