package basic

import (
	"fmt"
	"testing"
)

// 12. 使用复合字面值作为初值构造器

func TestLiteral(t *testing.T) {
	// ASCII 码，a代表数组的下标索引
	numbers := [256]int{'a': 8, 'b': 7, 'c': 4, 90: 0, 1: 1}
	fmt.Println(numbers)
}
