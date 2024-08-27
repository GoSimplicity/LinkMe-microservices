package publish

import (
	"context"
	"encoding/json"
	checkClient "github.com/GoSimplicity/LinkMe-microservices/api/check/v1"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"time"
)

type PublishPostEventConsumer struct {
	checkClient checkClient.CheckClient
	client      sarama.Client
	l           *zap.Logger
}

type Check struct {
	ID        int64     // 审核ID
	PostID    int64     // 帖子ID
	Content   string    // 审核内容
	Title     string    // 审核标签
	UserID    int64     // 提交审核的用户ID
	Status    uint8     // 审核状态
	Remark    string    // 审核备注
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

type consumerGroupHandler struct {
	consumer *PublishPostEventConsumer
}

// NewPublishPostEventConsumer 创建一个新的 PublishPostEventConsumer 实例
func NewPublishPostEventConsumer(checkClient checkClient.CheckClient, client sarama.Client, l *zap.Logger) *PublishPostEventConsumer {
	return &PublishPostEventConsumer{
		checkClient: checkClient,
		client:      client,
		l:           l,
	}
}

func (c *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// 处理每一条消息
		if err := c.consumer.processMessage(sess.Context(), msg); err != nil {
			c.consumer.l.Error("Failed to process message", zap.Error(err), zap.ByteString("message", msg.Value), zap.Int64("offset", msg.Offset))
		} else {
			// 如果消息处理成功，标记消息为已消费
			sess.MarkMessage(msg, "")
		}
	}
	return nil
}

// Start 启动消费者，并开始消费 Kafka 中的消息
func (p *PublishPostEventConsumer) Start(ctx context.Context) error {
	cg, err := sarama.NewConsumerGroupFromClient("publish_event", p.client)
	if err != nil {
		return err
	}

	p.l.Info("PublishConsumer started")

	go func() {
		for {
			// 开始消费指定的 Kafka 主题
			err := cg.Consume(ctx, []string{TopicPublishEvent}, &consumerGroupHandler{consumer: p})
			if err != nil {
				p.l.Error("Error occurred in consume loop", zap.Error(err))
				// 判断上下文是否已取消，若已取消则退出循环
				if ctx.Err() != nil {
					p.l.Info("Context canceled, stopping consumer")
					return
				}
				time.Sleep(time.Second) // 避免过于频繁的重试
			}
		}
	}()

	return nil
}

// processMessage 处理从 Kafka 消费的消息
func (p *PublishPostEventConsumer) processMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {
	var event PublishEvent
	// 将消息内容反序列化为 PublishEvent 结构体
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		return err
	}

	// 创建检查记录
	check := &checkClient.CreateCheckRequest{
		Content: event.Content,
		PostId:  event.PostId,
		Title:   event.Title,
		UserId:  event.UserID,
	}

	// 使用传递的上下文来调用 checkClient
	checkId, err := p.checkClient.CreateCheck(ctx, check)

	if err != nil {
		p.l.Error("Failed to create check", zap.Error(err))
		return err
	}

	p.l.Info("Successfully processed message", zap.String("check_id", checkId.String()))

	return nil
}
