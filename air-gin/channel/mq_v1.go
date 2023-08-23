package channel

import (
	"context"
	"errors"
	"sync"
)

type Msg struct {
	Content string
}

type BrokerService interface {
	Send(ctx context.Context, msg Msg) error
	Subscribe(ctx context.Context, capacity int) (chan Msg, error)
	Close() error
}

// Broker 实现一个多个消费组的模式
type Broker struct {
	mutex sync.Mutex
	chans []chan Msg // channel的数组，数组里面的元素为Msg
}

func (b *Broker) Send(ctx context.Context, msg Msg) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, ch := range b.chans {
		select {
		case ch <- msg:
		// 如果channel未满
		default:
			return errors.New("消息队列已满")
		}
	}

	return nil
}

func (b *Broker) Subscribe(ctx context.Context, capacity int) (chan Msg, error) {
	// 当一个消费者订阅capacity，那我们就建立一个这么一个大小的channel
	res := make(chan Msg, capacity)
	b.mutex.Lock()
	b.chans = append(b.chans, res)
	b.mutex.Unlock()

	return res, nil
}

// Close 这种方法体现的效率上吧，反正这里都加了锁的
func (b *Broker) Close() error {
	b.mutex.Lock()
	chans := b.chans
	b.chans = nil
	b.mutex.Unlock()

	// 避免重复
	for _, ch := range chans {
		close(ch)
	}

	return nil
}
