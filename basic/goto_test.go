package basic

import (
	"fmt"
	"testing"
	"time"
)

func TestGoTo(t *testing.T) {
	a := 10

Loop:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto Loop
		}
		t.Log(a)
		a = a + 1

	}

	fmt.Println("Hello, world!")
	a++
}

func TestTimer(t *testing.T) {
	go func() {
		time.AfterFunc(5*time.Second, func() {
			fmt.Println("demo timer")
		})
	}()

	time.Sleep(10 * time.Second)
}
