package data

import (
	"fmt"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/internal/conf"
	"github.com/IBM/sarama"
)

// NewSaramaClient 初始化Sarama客户端，用于连接到Kafka集群
func NewSaramaClient(c *conf.Data) (sarama.Client, error) {
	scfg := sarama.NewConfig()
	scfg.Producer.Return.Successes = true

	client, err := sarama.NewClient(c.Kafka.Addr, scfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Sarama client: %w", err)
	}

	return client, nil
}

// NewSyncProducer 使用已有的Sarama客户端初始化同步生产者
func NewSyncProducer(c sarama.Client) sarama.SyncProducer {
	p, err := sarama.NewSyncProducerFromClient(c)
	if err != nil {
		panic(err)
	}

	return p
}

//
//// NewConsumers 初始化并返回一个事件消费者
//func NewConsumers(publishConsumer *publish.PublishPostEventConsumer, syncConsumer *sync.SyncConsumer) []events.Consumer {
//	// 返回消费者切片
//	return []events.Consumer{publishConsumer, syncConsumer}
//}
