package fridayEngine

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/wire"
	"github.com/kuno989/friday/backend/pkg"
	"github.com/kuno989/friday/backend/schema/rabbitmq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"github.com/terra-farm/go-virtualbox"
	"os"
	"path/filepath"
	"time"
)

var (
	DefaultServerConfig = ServerConfig{
		Debug: true,
	}
	ServerProviderSet = wire.NewSet(NewServer, ProvideServerConfig)
)

type ServerConfig struct {
	Debug    bool   `mapstructure:"debug"`
	TempPath string `mapstructure:"volume"`
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
	body := bytes.ReplaceAll(msg.Body, []byte("NaN"), []byte("0"))
	var resp rabbitmq.ResponseObject
	if err := json.Unmarshal(body, &resp); err != nil {
		logrus.Error("failed to parse message body", err)
		if err := msg.Reject(false); err != nil {
			logrus.Error("failed to reject message", err)
		}
	}
	filePath := filepath.Join(s.Config.TempPath, resp.Sha256)
	file, err := os.Create(filePath)
	if err != nil {
		logrus.Error("failed creating file", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.minio.Download(ctx, resp.MinioObjectKey, file); err != nil {
		logrus.Error("failed downloading file", err)
	}
	file.Close()
	logrus.Debugf("file downloaded to %s", filePath)
	machine, err := virtualbox.GetMachine("win7")
	if err != nil {
		logrus.Error("can not find machine", err)
	}

	logrus.Debugf("%s sandbox found!", machine.Name)
	logrus.Debugf("cpu %v, memory %v", machine.CPUs, machine.Memory)
	return nil
}
