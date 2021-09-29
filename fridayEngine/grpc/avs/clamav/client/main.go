package client

import (
	"context"
	"github.com/kuno989/friday/fridayEngine/grpc/avs"
	pb "github.com/kuno989/friday/fridayEngine/grpc/avs/clamav/proto"
)

func GetVersion(client pb.ClamAVScannerClient) (*pb.VersionResponse, error) {
	version := &pb.VersionRequest{}
	return client.GetVersion(context.Background(), version)
}

func ScanFile(client pb.ClamAVScannerClient, path string) (avs.ScanResult, error) {
	scanFile := &pb.ScanFileRequest{Filepath: path}
	ctx, cancel := context.WithTimeout(context.Background(), avs.ScanTimeout)
	defer cancel()
	res, err := client.ScanFile(ctx, scanFile)
	if err != nil {
		return avs.ScanResult{}, err
	}
	return avs.ScanResult{
		Output:   res.Output,
		Infected: res.Infected,
	}, nil
}
