package client

import (
	"context"
	"github.com/kuno989/friday/fridayEngine/grpc/avs"
	pb "github.com/kuno989/friday/fridayEngine/grpc/avs/windefender/proto"
)

func GetVersion(client pb.WinDefenderScannerClient) (*pb.VersionResponse, error) {
	version := &pb.VersionRequest{}
	return client.GetVersion(context.Background(), version)
}

func ScanFile(client pb.WinDefenderScannerClient, path string) (avs.ScanResult, error) {
	scan := &pb.ScanFileRequest{Filepath: path}
	ctx, cancel := context.WithTimeout(context.Background(), avs.ScanTimeout)
	defer cancel()
	res, err := client.ScanFile(ctx, scan)
	if err != nil {
		return avs.ScanResult{}, err
	}
	return avs.ScanResult{
		Output:   res.Output,
		Infected: res.Infected,
	}, nil
}
