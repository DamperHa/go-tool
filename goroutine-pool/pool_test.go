package goroutine_pool

import (
	"testing"
	"time"
)

//
func TestName(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			t.Log("hello")
		}
	}
}
