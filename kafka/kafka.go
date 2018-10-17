package kafka

import (
	"github.com/Shopify/sarama"
	"os"
	"time"
	"fmt"
	"os/signal"
	"github.com/bsm/sarama-cluster"
	"strings"
)

// 同步生产者
func ProducerSync(topic, value string) {
	config := sarama.NewConfig()
	// 等待服务器所有副本保存成功后响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向分区(partition)发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// 设置kafka版本
	config.Version = sarama.V2_0_0_0

	fmt.Println("start make producer")
	producer, err := sarama.NewSyncProducer([]string{"10.221.164.60:9092"}, config)
	if err != nil {
		fmt.Printf("producer err: %s\n", err.Error())
		return
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
	}
	msg.Value = sarama.ByteEncoder(value)

	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("send message err: %s", err.Error())
		return
	}

	fmt.Printf("pid: %v, offset: %v", pid, offset)
}

// 异步生产者
func ProducerAsync(topic, value string) {
	config := sarama.NewConfig()
	// 等待服务器所有副本保存成功后响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向分区(partition)发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// 设置kafka版本
	config.Version = sarama.V2_0_0_0

	// kafka服务器，多个以英文逗号分隔
	brokers := "10.221.164.60:9092"

	fmt.Println("start make producer")
	producer, err := sarama.NewAsyncProducer(strings.Split(brokers, ","), config)
	if err != nil {
		fmt.Printf("producer err: %s\n", err.Error())
		return
	}
	defer producer.AsyncClose()

	msg := &sarama.ProducerMessage{
		Topic: topic,
	}
	msg.Value = sarama.ByteEncoder(value)

	Loop:
	for {
		producer.Input() <- msg
		select {
		case <-producer.Successes():
			fmt.Println("send message success")
			break Loop
		case err := <-producer.Errors():
			fmt.Printf("err: %s", err.Err.Error())
			return
		}
	}
}

// 消费者
func Consumer() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Version = sarama.MaxVersion
	consumer, err := sarama.NewConsumer([]string{"10.221.164.60:9092"}, config)
	if err != nil {
		fmt.Printf("consumer err:%s", err.Error())
		return
	}

	client, err := sarama.NewClient([]string{"10.221.164.60:9092"}, config)
	if err != nil {
		fmt.Printf("client err:%s", err.Error())
		return
	}
	defer client.Close()

	offsetManager, err := sarama.NewOffsetManagerFromClient("group111", client)
	if err != nil {
		fmt.Printf("offsetManager err:%s", err.Error())
		return
	}
	defer offsetManager.Close()

	partitionOffsetManager, err := offsetManager.ManagePartition("lang", 0)
	if err != nil {
		fmt.Printf("partitionOffsetManager err:%s", err.Error())
		return
	}
	defer partitionOffsetManager.Close()

	fmt.Println("consumer init success")

	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Printf("err:%s", err.Error())
		}
	}()

	topics, _ := consumer.Topics()
	fmt.Println(topics)
	partitions, _ := consumer.Partitions("lang")
	fmt.Println(partitions)
	nextOffset, _ := partitionOffsetManager.NextOffset()
	fmt.Println(nextOffset)

	partitionConsumer, err := consumer.ConsumePartition("lang", 0, nextOffset+1)
	if err != nil {
		fmt.Printf("partitionConsumer err:%s", err.Error())
		return
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			fmt.Printf("err:%s", err.Error())
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	fmt.Println("start consume really")

	ConsumerLoop:
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				fmt.Printf("message offset:%d, message:%s\n", msg.Offset, string(msg.Value))
				nextOffset, offsetString := partitionOffsetManager.NextOffset()
				fmt.Println(nextOffset+1, "...", offsetString)
				partitionOffsetManager.MarkOffset(nextOffset+1, "modified meatdata")
			case <-signals:
				break ConsumerLoop
			}
		}
}

// 消费者集群
func ConsumerCluster() {
	brokers := []string{"10.221.164.60:9092"}
	topics := []string{"lang"}
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Group.Return.Notifications = true

	consumer, err := cluster.NewConsumer(brokers, "consumer1", topics, config)
	if err != nil {
		fmt.Printf("consumer err: %s", err.Error())
		return
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		for err := range consumer.Errors() {
			fmt.Printf("err: %s", err.Error())
		}
	}()

	go func() {
		for ntf := range consumer.Notifications() {
			fmt.Printf("ntf: %+v\n", ntf)
		}
	}()

	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Printf("%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "")
			}
		case <-signals:
			return
		}
	}
}
