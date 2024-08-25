package events

import (
	"context"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/events/publish"
	"github.com/GoSimplicity/LinkMe-microservices/app/linkme-post/events/sync"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(sync.NewSyncConsumer, publish.NewSaramaSyncProducer, publish.NewPublishPostEventConsumer)

type Consumer interface {
	Start(ctx context.Context) error
}
