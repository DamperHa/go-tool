package channel

import (
	"context"
	"sync"
)

type BrokerV2 struct {
	mutex     sync.Mutex
	consumers []func(msg Msg)
}

func (b *BrokerV2) Send(ctx context.Context, msg Msg) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for _, consumer := range b.consumers {
		consumer(msg)
	}

	return nil
}

// Subscribe 当存在多个协程修改某个对象时，加锁就对了
func (b *BrokerV2) Subscribe(consumer func(msg Msg)) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.consumers = append(b.consumers, consumer)

	return nil
}
