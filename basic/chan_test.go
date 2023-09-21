package basic

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 6)
	ch <- 2
	ch <- 3

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

type T struct{}

func spawn(f func()) chan T {
	ch := make(chan T)
	go func() {

		f()
	}()

	return ch
}

// 创建模型
// 两个goroutine可以通过channel进行通信
func TestCreatePattern(t *testing.T) {
	ch := spawn(func() {})
	fmt.Println(ch)
}

// Join模型，goroutine创建者需要等待新的goroutine结束，这种起名为"join模式"
func worker(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	interval, ok := args[0].(int)
	if !ok {
		return
	}

	time.Sleep(time.Second * (time.Duration(interval)))
}

func spawn2(f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})
	go func() {
		f(args...)
		c <- struct{}{}
	}()

	return c
}

func TestJoin(t *testing.T) {
	done := spawn2(worker, 5)
	println("spwan a worker goroutine")

	// 阻塞在这，等待新创建的goroutine完成
	<-done
}

// 如果，我还需要获取新创建协程的退出状态呢
var OK = errors.New("ok")

func worker3(args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("invalid args")
	}

	interval, ok := args[0].(int)
	if !ok {
		return errors.New("invalid interval arg")
	}

	time.Sleep(time.Second * (time.Duration(interval)))

	return OK
}

func spawn3(f func(args ...interface{}) error, args ...interface{}) chan error {
	c := make(chan error)
	go func() {
		c <- f(args...)
	}()

	return c
}

func TestJoinError(t *testing.T) {
	done := spawn3(worker3, 5)
	println("spwan a worker goroutine")

	// 阻塞在这，等待新创建的goroutine完成
	<-done
}

