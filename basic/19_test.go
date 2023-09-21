package basic

import "testing"

// 这还能这样写呢？
func TestCodeBlock(t *testing.T) {
	if a := 1; false {

	} else if b := 2; false {

	} else if c := 3; false {

	} else {
		println(a, b, c) // 输出 1 2 3
	}
}
