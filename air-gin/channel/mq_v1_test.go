package channel

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMqSendSubscribe(t *testing.T) {
	b := &Broker{}
	ctx := context.Background()

	// 模拟发送者
	go func() {
		for {
			err := b.Send(ctx, Msg{Content: time.Now().String()})
			if err != nil {
				t.Log(err)
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("消费者 %d", i)
		go func() {
			defer wg.Done()
			resChan, err := b.Subscribe(ctx, 100)
			if err != nil {
				t.Log(err)
			}

			// 针对于channel，数组，这几种写法
			for msg := range resChan {
				fmt.Println(name, msg)
			}
		}()
	}

	wg.Wait()
}

func TestBrokerV2(t *testing.T) {
	broker := &BrokerV2{}
	ctx := context.Background()

	go func() {
		for {
			err := broker.Send(ctx, Msg{time.Now().String()})
			if err != nil {
				t.Log(err)
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("消费者 %d", i)
		err := broker.Subscribe(func(msg Msg) {
			fmt.Println(name, msg.Content)
		})
		if err != nil {
			t.Log(err)
		}
	}

	select {}
}
