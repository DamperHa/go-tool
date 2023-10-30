package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"math/rand"
	"strings"
	"testing"
)

const (
	broker = "127.0.0.1:9093"
	topic  = "MyTopic"
)

// 异步发送消息
func TestProducer(t *testing.T) {
	brokerList := strings.Split(broker, ",")

	producer, err := newProducer(brokerList)
	if err != nil {
		log.Fatal(err)
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder("random_number"),
		Value: sarama.StringEncoder(fmt.Sprintf("%d", rand.Intn(1000))),
	}

	producer.Input() <- &msg
	successMsg := <-producer.Successes()
	log.Println("Successful to write message, offset:", successMsg.Offset)

	err = producer.Close()
	if err != nil {
		log.Fatal("Failed to close producer:", err)
	}
}

// 同步发送消息嗯

func newProducer(brokerList []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return nil, fmt.Errorf("starting Sarama producer:%w", err)
	}

	go func() {
		for err := range producer.Errors() {
			log.Println("Failed to write message:", err)
		}
	}()

	return producer, nil
}

func TestSync(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，默认设置8个分区
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		fmt.Println("producer close err, ", err)
		return
	}
	defer client.Close()

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder("Hello World!")

	pid, offset, err := client.SendMessage(msg)

	if err != nil {
		fmt.Println("send message failed, ", err)
		return
	}
	fmt.Printf("分区ID:%v, offset:%v \n", pid, offset)
}
