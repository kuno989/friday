package pkg

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/kuno989/friday/backend/schema"
	mongoSchema "github.com/kuno989/friday/backend/schema/mongo"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	DefaultMongoConfig = MongoConfig{
		URI:            "mongodb://localhost",
		DB:             "",
		FileCollection: "",
		BcryptCost:     10,
	}
	MongoProviderSet = wire.NewSet(NewMongo, ProvideMongoConfig)
)

type MongoConfig struct {
	URI            string `mapstructure:"uri"`
	DB             string `mapstructure:"db"`
	FileCollection string `mapstructure:"file_collection"`
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

func NewMongo(ctx context.Context, cfg MongoConfig) (*Mongo, func(), error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, nil, err
	}
	ctxTime, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := client.Connect(ctxTime); err != nil {
		return nil, nil, fmt.Errorf("connect mongo: %w", err)
	}
	cleanup := func() {
		_ = client.Disconnect(ctx)
	}
	return &Mongo{
		Config: cfg,
		client: client,
	}, cleanup, nil
}

func (m *Mongo) FileSearch(ctx context.Context, sha256 string) (schema.File, error) {
	coll := m.client.Database(m.Config.DB).Collection(m.Config.FileCollection)
	var file schema.File
	err := coll.FindOne(ctx, bson.M{mongoSchema.FileSha256Key: sha256}).Decode(&file)
	return file, err
}

func (m *Mongo) CreateFile(ctx context.Context, uploadFile schema.File) (schema.File, error) {
	coll := m.client.Database(m.Config.DB).Collection(m.Config.FileCollection)
	_, err := coll.InsertOne(ctx, uploadFile)
	if err != nil {
		return schema.File{}, nil
	}
	return uploadFile, nil
}

func (m *Mongo) UpdateFile(ctx context.Context, uploadFile schema.File) (schema.File, error) {
	coll := m.client.Database(m.Config.DB).Collection(m.Config.FileCollection)
	_, err := coll.UpdateOne(ctx, bson.M{"sha256": uploadFile.Sha256}, bson.M{
		"$set": uploadFile,
	})
	if err != nil {
		return schema.File{}, nil
	}
	return uploadFile, nil
}
