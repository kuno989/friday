package fridayEngine

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/kuno989/friday/backend/pkg"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var (
	DefaultServerConfig = ServerConfig{
		Debug: true,
	}
	ServerProviderSet = wire.NewSet(NewServer, ProvideServerConfig)
)

type ServerConfig struct {
	Debug bool `mapstructure:"debug"`
}

func ProvideServerConfig(cfg *viper.Viper) (ServerConfig, error) {
	sc := DefaultServerConfig
	err := cfg.Unmarshal(&sc)
	return sc, err
}

type Server struct {
	Config ServerConfig
	Rb     *pkg.RabbitMq
	ms     *pkg.Mongo
	minio  *pkg.Minio
}

func NewServer(cfg ServerConfig, ms *pkg.Mongo, rb *pkg.RabbitMq, minio *pkg.Minio) *Server {
	s := &Server{
		Config: cfg,
		Rb:     rb,
		ms:     ms,
		minio:  minio,
	}
	return s
}

func (s *Server) AmqpHandler(msg amqp.Delivery) error {
	if len(msg.Body) == 0 {
		return errors.New("Delivery Body is length 0")
	}
	msgBody := string(msg.Body)
	fmt.Println(msgBody)
	return nil
}
