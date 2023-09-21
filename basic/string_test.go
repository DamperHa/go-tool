package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFor(t *testing.T) {
	str := "fan zhi hao"

	for i := 0; i < len(str); i++ {
		fmt.Printf("编号：i:[%d],  [%v]", i, str[i])
	}
}

func TestCharacter(t *testing.T) {
	c := 'a'

	fmt.Println(reflect.TypeOf(c))
}

func TestLen(t *testing.T) {
	s := "这o"

	// output4: 这：3个字节表示，o：一个字节表示
	fmt.Println(len(s)) // 返回的是字节数，goang使用UTF-8的变长编码

	// 将字符串转换为[]rune
	fmt.Println(len([]rune(s)))
}

func TestAccess(t *testing.T) {
	str := "xiao范"

	s := str[1:2]
	fmt.Println(reflect.TypeOf(s))      // string类型
	fmt.Println(reflect.TypeOf(str[1])) // unit8

	for i := 0; i < len(str); i++ {
		fmt.Printf("编号：i:[%d],  [%#U] \n", i, str[i])
	}

	for i, r := range str {
		fmt.Printf("编号：i:[%d],  [%#U] \n", i, r)
	}
}

// 字符串拼接的几种方式
