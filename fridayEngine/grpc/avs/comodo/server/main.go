package main

import (
	"context"
	"github.com/kuno989/friday/fridayEngine/avs/comodo"
	"github.com/kuno989/friday/fridayEngine/grpc/avs"
	pb "github.com/kuno989/friday/fridayEngine/grpc/avs/comodo/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/grpclog"
)

type server struct {
	log *zap.Logger
	pb.UnsafeComodoScannerServer
}

func (s *server) GetVersion(ctx context.Context, in *pb.VersionRequest) (*pb.VersionResponse, error) {
	//version, err := comodo.GetVersion()
	return &pb.VersionResponse{Version: "version"}, nil
}

func (s *server) ScanFile(ctx context.Context, in *pb.ScanFileRequest) (*pb.ScanResponse, error) {
	s.log.Info("스캔:", zap.String("path", in.Filepath))
	res, err := comodo.ScanFile(in.Filepath)
	return &pb.ScanResponse{
		Infected: res.Infected,
		Output:   res.Output}, err
}

func main() {
	log := avs.LoggingSetup()
	lis, err := avs.MakeListener()
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	s := avs.NewServer()
	pb.RegisterComodoScannerServer(s, &server{log: log})
	log.Info("Comodo server running ...")
	err = avs.Serve(s, lis)
	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
	}
}
