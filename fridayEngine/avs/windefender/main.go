package windefender

import (
	"github.com/kuno989/friday/fridayEngine/utils"
	"os"
	"path"
	"strings"
)

const (
	loadlibraryPath = "/opt/windows-defender/"
	mpclient        = "./mpclient"
	mpenginedll     = "/engine/mpengine.dll"
)

type Result struct {
	Infected bool   `json:"infected"`
	Output   string `json:"output"`
}

func ScanFile(path string) (Result, error) {
	dir, err := utils.Getwd()
	if err != nil {
		return Result{}, err
	}
	if err := os.Chdir(loadlibraryPath); err != nil {
		return Result{}, err
	}
	defer os.Chdir(dir)
	output, err := utils.CMD(mpclient, path)
	if err != nil {
		return Result{}, err
	}

	res := Result{}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if !strings.Contains(line, "EngineScanCallback(): Threat ") {
			continue
		}

		detection := strings.TrimPrefix(line, "EngineScanCallback(): Threat ")
		res.Output = strings.TrimSuffix(detection, " identified.")
		res.Infected = true
		break
	}

	return res, nil
}

func GetVersion() (string, error) {
	mpenginedll := path.Join(loadlibraryPath, mpenginedll)
	versionOut, err := utils.CMD("exiftool", "-ProductVersion",
		mpenginedll)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(strings.Split(versionOut, ":")[1]), nil
}
