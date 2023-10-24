package panic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 如何处理panic呢
func TestThrow(t *testing.T) {
	var lock sync.Mutex
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("fix")
		}
	}()

	// lock.Lock()
	lock.Unlock()

	fmt.Println("hello world")
}

func TestMap(t *testing.T) {
	mp := make(map[int]int, 0)

	for i := 0; i < 10; i++ {
		go func() {
			mp[i] = 1
		}()
	}

	time.Sleep(time.Second * 10)
}
