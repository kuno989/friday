package fridayEngine

import (
	"fmt"
	"github.com/kuno989/friday/fridayEngine/grpc/avs"
	clamav_client "github.com/kuno989/friday/fridayEngine/grpc/avs/clamav/client"
	clamav_api "github.com/kuno989/friday/fridayEngine/grpc/avs/clamav/proto"
	comodo_client "github.com/kuno989/friday/fridayEngine/grpc/avs/comodo/client"
	comodo_api "github.com/kuno989/friday/fridayEngine/grpc/avs/comodo/proto"
	drweb_client "github.com/kuno989/friday/fridayEngine/grpc/avs/drweb/client"
	drweb_api "github.com/kuno989/friday/fridayEngine/grpc/avs/drweb/proto"
	windefender_client "github.com/kuno989/friday/fridayEngine/grpc/avs/windefender/client"
	windefender_api "github.com/kuno989/friday/fridayEngine/grpc/avs/windefender/proto"
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
	connect, err := avs.GetClientConn(s.Config.AVConfig[engine], engine)
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
	case "windefender":
		res, err = windefender_client.ScanFile(windefender_api.NewWinDefenderScannerClient(connect), path)
	case "drweb":
		res, err = drweb_client.ScanFile(drweb_api.NewDrWebScannerClient(connect), path)
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
	winChannel := make(chan avs.ScanResult)
	drwebChannel := make(chan avs.ScanResult)

	go s.avScan("comodo", path, comodoChannel)
	go s.avScan("clamav", path, clamavChannel)
	go s.avScan("windefender", path, winChannel)
	go s.avScan("drweb", path, drwebChannel)

	avScanResults := map[string]interface{}{}
	engineCounts := 4
	count := 0
	for {
		select {
		case comodoResponse := <-comodoChannel:
			avScanResults["comodo"] = comodoResponse
			count++
		case clamavResponse := <-clamavChannel:
			avScanResults["clamav"] = clamavResponse
			count++
		case windefenderResponse := <-winChannel:
			avScanResults["windefender"] = windefenderResponse
			count++
		case drwebResponse := <-drwebChannel:
			avScanResults["drweb"] = drwebResponse
			count++
		}
		if count == engineCounts {
			break
		}
	}

	return avScanResults
}
