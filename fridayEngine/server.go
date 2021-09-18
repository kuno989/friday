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
	body := bytes.ReplaceAll(msg.Body, []byte("NaN"), []byte("0"))
	var resp rabbitmq.ResponseObject
	if err := json.Unmarshal(body, &resp); err != nil {
		return errors.New("ailed to parse message body")
		if err := msg.Reject(false); err != nil {
			logrus.Errorf("failed to reject message %s", err)
		}
	}
	filePath := filepath.Join(s.Config.TempPath, resp.Sha256)
	file, err := os.Create(filePath)
	if err != nil {
		logrus.Errorf("failed creating file %s", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.minio.Download(ctx, resp.MinioObjectKey, file); err != nil {
		logrus.Errorf("failed downloading file %s", err)
	}
	file.Close()
	logrus.Infof("file downloaded to %s", filePath)
	s.defaultScan(filePath)

	//machine, err := virtualbox.GetMachine("win7")
	//if err != nil {
	//	logrus.Errorf("can not find machine %s", err)
	//}
	//
	//logrus.Infof("%s sandbox found", machine.Name)
	//logrus.Infof("cpu %v, memory %v", machine.CPUs, machine.Memory)
	//
	//if err := machine.Start(); err != nil {
	//	logrus.Errorf("machine start failure %s", err)
	//}
	//logrus.Infof("%s sandbox start", machine.Name)

	//if err := machine.Save(); err != nil {
	//	logrus.Errorf("machine save failure %s", err)
	//}
	//logrus.Infof("%s sandbox saved", machine.Name)

	return nil
}
