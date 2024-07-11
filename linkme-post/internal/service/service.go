package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet4 = wire.NewSet(NewPostService)
