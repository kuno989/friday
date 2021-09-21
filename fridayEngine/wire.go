//go:build wireinject
// +build wireinject

package fridayEngine

import (
	"context"

	"github.com/google/wire"
	"github.com/spf13/viper"

	"github.com/kuno989/friday/backend/pkg"
)

func InitializeServer(ctx context.Context, cfg *viper.Viper) (*Server, func(), error) {
	panic(wire.Build(ServerProviderSet, pkg.MongoProviderSet, pkg.RabbitmqProviderSet, pkg.MinioProviderSet, pkg.YaraProviderSet))
}
