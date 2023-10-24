package _select

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	default:
	}
}

func TestLoopSelect(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		// 每隔1s发送一条消息到channel中
		for range time.Tick(1 * time.Second) {
			ch1 <- 1
		}
	}()

	for {
		select {
		case <-ch1:
			fmt.Println("ch1")
		case <-ch2:
			fmt.Println("ch2")
		}
	}
}
