package basic

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestPointer(t *testing.T) {
	res := &Person{}

	res.name = "zhangsan"
	res.age = 18

	fmt.Println(res)

	res = &Person{
		name: "lisi",
		age:  20,
	}

	fmt.Println(res)

	// 字符串的复制
	changePointer(res)
	fmt.Println(res)
}

func changePointer(arg *Person) {

	arg = &Person{
		name: "111",
		age:  12,
	}
}
