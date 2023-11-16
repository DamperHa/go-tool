package basic

import (
	"context"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			println("ok")
			time.Sleep(1 * time.Second)
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	time.Sleep(5 * time.Second)
}
