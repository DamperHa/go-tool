package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	consumerGroupID = "sarama_consumer"
)

// ConsumerGroupHandler represents the sarama consumer group
type ConsumerGroupHandler struct{}

// Setup is run before consumer start consuming, is normally used to setup things such as database connections
func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error { return nil }

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages(), here is supposed to be what you want to
// do with the message. In this example the message will be logged with the topic name, partition and message value.
func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d message: %v\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		time.Sleep(2 * time.Second)
		sess.MarkMessage(msg, "")
		sess.Commit()
	}
	return nil
}

func TestConsumer(t *testing.T) {
	config := sarama.NewConfig()
	config.ClientID = "1"
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
	config.Version = sarama.V2_5_0_0
	config.Consumer.Offsets.AutoCommit.Enable = false

	brokerList := strings.Split(broker, ",")

	client, err := sarama.NewClient(brokerList, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = client.Close()
	}()

	group, err := sarama.NewConsumerGroupFromClient(consumerGroupID, client)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = group.Close()
	}()

	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	ctx := context.Background()
	for {
		handler := ConsumerGroupHandler{}
		err := group.Consume(ctx, []string{topic}, handler)
		if err != nil {
			panic(err)
		}
	}
}

func TestSingleConsumer(t *testing.T) {
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}

			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}

	wg.Wait()
}
