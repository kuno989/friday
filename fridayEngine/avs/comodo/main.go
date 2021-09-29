package comodo

import (
	"fmt"
	"github.com/kuno989/friday/fridayEngine/utils"
	"strings"
)

const (
	cmdscan   = "/opt/COMODO/cmdscan"
	comodover = "/opt/COMODO/cavver.dat"
)

type Result struct {
	Infected bool   `json:"infected"`
	Output   string `json:"output"`
}

func ScanFile(path string) (Result, error) {
	cmd, err := utils.CMD(cmdscan, "-v", "-s", path)
	if err != nil {
		return Result{}, err
	}
	lines := strings.Split(cmd, "\n")
	if len(lines) < 2 {
		errUnexpectedOutput := fmt.Errorf("unexpected output: %s", cmd)
		return Result{}, errUnexpectedOutput
	}
	if strings.HasSuffix(lines[1], "---> Not Virus") {
		return Result{}, nil
	}
	res := Result{}
	detection := strings.Split(lines[1], "Malware Name is ")
	res.Output = detection[len(detection)-1]
	res.Infected = true
	return res, nil
}

func GetVersion() (string, error) {
	version, err := utils.ReadAll(comodover)
	if err != nil {
		return "", err
	}
	return string(version), nil
}
