package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet3 = wire.NewSet(NewPostUsecase)
