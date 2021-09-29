package fridayEngine

import (
	"fmt"
	"github.com/kuno989/friday/fridayEngine/grpc/avs"
	clamav_client "github.com/kuno989/friday/fridayEngine/grpc/avs/clamav/client"
	clamav_api "github.com/kuno989/friday/fridayEngine/grpc/avs/clamav/proto"
	comodo_client "github.com/kuno989/friday/fridayEngine/grpc/avs/comodo/client"
	comodo_api "github.com/kuno989/friday/fridayEngine/grpc/avs/comodo/proto"
	"github.com/kuno989/friday/fridayEngine/utils"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

func (s *Server) avScan(engine, path string, c chan avs.ScanResult) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("panic av scan %v", debug.Stack())
		}
	}()
	connect, err := avs.GetClientConn(s.Config.AVConfig[engine])
	if err != nil {
		logrus.Errorf("grpc client connect [%s]: %v", engine, err)
		c <- avs.ScanResult{}
		return
	}
	defer connect.Close()

	copyPath := fmt.Sprintf("%s-%s", path, engine)
	err = utils.CopyFile(path, copyPath)
	if err != nil {
		logrus.Errorf("Failed to copy %s", engine)
		c <- avs.ScanResult{}
		return
	}
	path = copyPath
	res := avs.ScanResult{}

	switch engine {
	case "comodo":
		res, err = comodo_client.ScanFile(comodo_api.NewComodoScannerClient(connect), path)
	case "clamav":
		res, err = clamav_client.ScanFile(clamav_api.NewClamAVScannerClient(connect), path)
	}

	if err != nil {
		logrus.Errorf("Failed to scan file %s : %v", engine, err)
	}

	c <- avs.ScanResult{Output: res.Output, Infected: res.Infected}

	if utils.Exists(copyPath) {
		if err = utils.DeleteFile(copyPath); err != nil {
			logrus.Errorf("Failed to delete %s", copyPath)
		}
	}
}

func (s *Server) parallelAvScan(path string) map[string]interface{} {

	comodoChannel := make(chan avs.ScanResult)
	clamavChannel := make(chan avs.ScanResult)

	go s.avScan("comodo", path, comodoChannel)
	go s.avScan("clamav", path, clamavChannel)

	avScanResults := map[string]interface{}{}
	engineCounts := 2
	count := 0
	for {
		select {
		case comodoResponse := <-comodoChannel:
			avScanResults["comodo"] = comodoResponse
			count++
		case clamavResponse := <-clamavChannel:
			avScanResults["clamav"] = clamavResponse
			count++
		}
		if count == engineCounts {
			break
		}
	}

	return avScanResults
}
