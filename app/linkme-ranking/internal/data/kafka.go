package data

import (
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-ranking/internal/conf"
	"github.com/IBM/sarama"
)

// NewSaramaClient 初始化一个新的 Sarama 客户端
func NewSaramaClient(c *conf.Service) (sarama.Client, error) {
	scfg := sarama.NewConfig()
	scfg.Producer.Return.Successes = true // 配置生产者需要返回确认成功的消息
	client, err := sarama.NewClient(c.Kafka.Addr, scfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewSyncProducer 初始化一个新的 Sarama 同步生产者
func NewSyncProducer(client sarama.Client) (sarama.SyncProducer, error) {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	return producer, nil
}
