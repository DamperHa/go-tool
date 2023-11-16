package error

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// Unwrap() error
// Unwrap() []error

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func TestName(t *testing.T) {
	if err := oops(); err != nil {

		// 这里并不是直接打印某个结构体，而是执行对应的error方法
		fmt.Println(err)
	}

}

var (
	PreErr = errors.New("origin error")
)

// 下面例子介绍了Wrap以及UnWrap
func TestWrapUnWrap(t *testing.T) {
	e2 := fmt.Errorf("这是第二个错误：%w", PreErr)
	e3 := fmt.Errorf("这是第三个错误：%w", e2)

	fmt.Println(fmt.Errorf("这是第四个错误：%w", e3))

	// 测试UnWrap
	unwrapE2 := errors.Unwrap(e3)
	fmt.Println(unwrapE2)

	origin := errors.Unwrap(errors.Unwrap(e3))

	// 通过一层一层的Unwrap，最终获取的是原始的错误
	if origin == PreErr {
		fmt.Println("true")
	}

}

//
