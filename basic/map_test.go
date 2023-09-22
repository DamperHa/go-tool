package basic

import (
	"fmt"
	"testing"
	"time"
)

// 什么的panic，能够被recover住呢？
func TestMap(t *testing.T) {
	mp := make(map[int]int)

	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()

			mp[1] = 2
		}()
	}

	time.Sleep(time.Minute)
}
