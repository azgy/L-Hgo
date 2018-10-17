package kafka

import (
	"testing"
	"time"
)

func TestProducerSync(t *testing.T) {
	ProducerSync("lang", "this is a message " + time.Now().Format("15:05:06"))
}

func TestProducerAsync(t *testing.T) {
	ProducerAsync("lang", "this is a message " + time.Now().Format("15:05:06"))
}

func TestConsumer(t *testing.T) {
	Consumer()
}

func TestConsumerCluster(t *testing.T) {
	ConsumerCluster()
}
