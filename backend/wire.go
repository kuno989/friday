//+build wireinject

package backend

import (
	"context"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

func InitializeServer(ctx context.Context, cfg *viper.Viper) (*Server, func(), error) {
	panic(wire.Build(ServerProviderSet))
}
