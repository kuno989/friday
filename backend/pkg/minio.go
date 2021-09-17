package pkg

import (
	"github.com/google/wire"
	miniogo "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)


var (
	DefaultMinioConfig = MinioConfig{
		URI: "localhost:9000",
		AccessKey: "",
		SecretKey: "",
	}
	MinioProviderSet = wire.NewSet(NewMinio, ProvideMinioConfig)
)
type MinioConfig struct {
	URI string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
}
func ProvideMinioConfig(cfg *viper.Viper) (MinioConfig, error){
	minio := DefaultMinioConfig
	err := cfg.UnmarshalKey("minio", &minio)
	return minio, err
}

type Minio struct {
	Config MinioConfig
	Client *miniogo.Client
}

func NewMinio(cfg MinioConfig) (*Minio, error) {
	client, err := miniogo.New(cfg.URI, &miniogo.Options{
		Creds: credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
	})
	if err != nil {
		return nil, err
	}
	return &Minio{
		Config: cfg,
		Client: client,
	}, nil
}