package basic

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func childRoutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Child goroutine received cancellation signal.")
			return // 在收到取消信号后退出协程
		default:
			// 模拟子协程的工作
			fmt.Println("Child goroutine is working...")
			time.Sleep(10 * time.Second)
			fmt.Println("Child goroutine is done...")

		}
	}
}

func TestDemo(t *testing.T) {
	// 创建一个父context，它可以取消子协程
	parentCtx, cancel := context.WithCancel(context.Background())

	// 启动子协程来执行某项任务
	go childRoutine(parentCtx)

	// 模拟一段时间后取消任务
	time.Sleep(3 * time.Second)
	fmt.Println("Cancelling the task...")
	cancel() // 取消任务

	// 等待一段时间，以确保子协程有足够的时间响应取消操作
	time.Sleep(20 * time.Second)
	fmt.Println("Main goroutine exits.")
}

func TestTimeout(t *testing.T) {
	// 创建一个父context，它可以取消子协程
	parentCtx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*3))

	// 启动子协程来执行某项任务
	go childRoutine(parentCtx)

	// 模拟一段时间后取消任务
	time.Sleep(6 * time.Second)
	fmt.Println("Cancelling the task...")
	cancel() // 取消任务

	// 等待一段时间，以确保子协程有足够的时间响应取消操作
	time.Sleep(20 * time.Second)
	fmt.Println("Main goroutine exits.")
}
