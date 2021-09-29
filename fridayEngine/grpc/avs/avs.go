package avs

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type Scanner interface {
	Scan()
}

func LoggingSetup() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	config.OutputPaths = []string{"stdout"}
	logger, _ := config.Build()
	return logger
}

type ScanResult struct {
	Output   string `json:"output"`
	Infected bool   `json:"infected"`
}

const (
	maxMsgSize  = 1024 * 1024 * 64
	port        = ":50051"
	ScanTimeout = 10 * time.Second
)

func DefaultServerOpts() []grpc.ServerOption {
	return []grpc.ServerOption{grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
	}
}

func NewServer(opts ...grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(append(DefaultServerOpts(), opts...)...)
}

func MakeListener() (net.Listener, error) {
	lis, err := net.Listen("tcp", port)
	return lis, err
}

func Serve(s *grpc.Server, lis net.Listener) error {
	reflection.Register(s)
	return s.Serve(lis)
}

func GetClientConn(address, engine string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		address, []grpc.DialOption{grpc.WithInsecure()}...)
	logrus.Infof("%s connected on %s", engine, port)
	return conn, err
}
