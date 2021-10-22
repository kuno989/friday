package fridayEngine

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/kuno989/friday/backend/pkg"
	"github.com/kuno989/friday/backend/schema"
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
	Debug         bool              `mapstructure:"debug"`
	TempPath      string            `mapstructure:"volume"`
	URI           string            `mapstructure:"uri"`
	AgentURI      string            `mapstructure:"agent_uri"`
	WebserverPort string            `mapstructure:"webserver_port"`
	AgentPort     string            `mapstructure:"agent_port"`
	VmName        string            `mapstructure:"vm_name"`
	AVConfig      map[string]string `mapstructure:"av"`
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
	yara   *pkg.Yara
}

const (
	queued       = iota
	processing   = iota
	finished     = iota
	vmProcessing = iota
)

func NewServer(cfg ServerConfig, ms *pkg.Mongo, rb *pkg.RabbitMq, minio *pkg.Minio, yara *pkg.Yara) *Server {
	s := &Server{
		Config: cfg,
		Rb:     rb,
		ms:     ms,
		minio:  minio,
		yara:   yara,
	}
	return s
}

func (s *Server) AmqpHandler(msg amqp.Delivery) error {
	body := bytes.ReplaceAll(msg.Body, []byte("NaN"), []byte("0"))
	var resp rabbitmq.ResponseObject
	if err := json.Unmarshal(body, &resp); err != nil {
		return errors.New("ailed to parse message body")
		if err := msg.Reject(false); err != nil {
			logrus.Errorf("failed to reject message %v", err)
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

	res := schema.Result{}
	res.Status = processing
	var buff []byte

	// static 분석 작업 전 업데이트 [status 변경]
	if buff, err = json.Marshal(res); err != nil {
		logrus.Errorf("Failed to json marshall object: %v ", err)
	}
	s.updateDocument(resp.Sha256, buff)

	s.defaultScan(filePath, &res)
	logrus.Info("Scan finished")

	// static 분석작업 완료후 업데이트 [status 및 분석결과 업데이트]

	res.Status = finished
	now := time.Now().UTC()
	res.LastScanned = &now

	if buff, err = json.Marshal(res); err != nil {
		logrus.Errorf("Failed to json marshall object: %v ", err)
	}
	s.updateDocument(resp.Sha256, buff)

	ch, err := s.Rb.Channel()
	if err != nil {
		logrus.Error("rabbitmq channel error", err)
	}
	defer ch.Close()

	// friday connect 큐에 등록
	// friday connect 에서 작업 실행
	q, err := ch.QueueDeclare(s.Rb.Config.VmQueue, false, false, false, false, nil)
	if err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        []byte(fmt.Sprintf(`{"minio_object_key":"%s", "sha256":"%s", "file_type":"%s"}`, resp.MinioObjectKey, resp.Sha256, res.Type)),
	}); err != nil {
		logrus.Info("vm 작업 큐 등록 완료")
	}

	return nil
}
