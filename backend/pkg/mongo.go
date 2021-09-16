package pkg

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	DefaultMongoConfig = MongoConfig{
		URI: "mongodb://localhost",
		DB: "friday",
		BcryptCost: 10,
	}
	MongoProviderSet = wire.NewSet(NewMongo, ProvideMongoConfig)
)
type MongoConfig struct {
	URI string `mapstructure:"uri"`
	DB string  `mapstructure:"db"`
	//UserCollection string
	//DatabaseCollection string
	BcryptCost int
}

func ProvideMongoConfig(cfg *viper.Viper) (MongoConfig, error) {
	mc := DefaultMongoConfig
	err := cfg.UnmarshalKey("mongo", &mc)
	return mc, err
}

type Mongo struct {
	Config MongoConfig
	client *mongo.Client
}

func NewMongo(ctx context.Context, cfg MongoConfig) (*Mongo, func(), error){
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, nil, err
	}
	ctxTime, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := client.Connect(ctxTime); err != nil {
		return nil, nil, fmt.Errorf("connect: %w", err)
	}
	cleanup := func() {
		_ = client.Disconnect(ctx)
	}
	return &Mongo{
		Config: cfg,
		client: client,
	}, cleanup, nil
}
func (m *Mongo) EnsureIndexes(ctx context.Context) error {
	return nil
}