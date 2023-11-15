package singleflight

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"testing"
	"time"
)

func TestSingleFlight(t *testing.T) {
	// 创建一个 SingleFlight Group
	var group singleflight.Group

	// 创建一个等待组，以便等待所有并发请求完成
	var wg sync.WaitGroup

	// 模拟并发请求
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			// 调用 Do 方法，确保只有一个 goroutine 执行这个函数
			val, err, _ := group.Do("example-key", func() (interface{}, error) {

				// 模拟一些耗时的操作
				time.Sleep(1 * time.Second)
				return fmt.Sprintf("Result for request %d", id), nil
			})

			// 处理结果
			if err != nil {
				fmt.Printf("Error for request %d: %v\n", id, err)
			} else {
				fmt.Printf("Result for request %d: %v\n", id, val)
			}
		}(i)
	}

	// 等待所有并发请求完成
	wg.Wait()
}

func doSomeThingWithRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered:", err)
		}
	}()

	// 模拟发生panic的情况
	panic("Something went wrong!")
}

func goWithPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered:", err)
		}
	}()

	fmt.Println("pre panic")

	// 这个协程里面的panic就recover不住了，所以会导致
	go panic("running panic")
	select {}
}

func TestPanic(t *testing.T) {
	go doSomeThingWithRecover()

	fmt.Println("Main goroutine continues.....")

	time.Sleep(10 * time.Second)
}

func TestGoPanic(t *testing.T) {
	go goWithPanic()

	fmt.Println("Main goroutine continues.....")

	time.Sleep(10 * time.Second)
}
