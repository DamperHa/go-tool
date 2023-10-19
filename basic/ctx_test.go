package basic

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
超时控制
*/

func hardWork(job interface{}) error {
	time.Sleep(time.Minute)
	return nil
}

func requestWorkV1(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	done := make(chan error)

	// done 退出以后，没有接受者，会导致协程阻塞
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// 可以做到超时控制，但是会出现协程泄露的情况
func TestV1(t *testing.T) {
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()

	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWorkV1(context.Background(), "any")
		}()
	}

	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))

	time.Sleep(time.Minute * 2)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}

func requestWorkV2(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	done := make(chan error, 1)

	// done 退出以后，没有接受者，会导致协程阻塞
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// 解决了协程泄露的问题
func TestV2(t *testing.T) {
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()

	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWorkV1(context.Background(), "any")
		}()
	}

	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))

	time.Sleep(time.Minute * 2)
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}
