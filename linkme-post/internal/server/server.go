package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet1 = wire.NewSet(NewGRPCServer, NewHTTPServer)
