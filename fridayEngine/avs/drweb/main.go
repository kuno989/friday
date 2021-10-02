package drweb

import (
	"github.com/kuno989/friday/fridayEngine/utils"
	"regexp"
	"strings"
)

const (
	cmd      = "/opt/drweb.com/bin/drweb-ctl"
	regexStr = "infected with (.*)"
)

type Result struct {
	Infected bool   `json:"infected"`
	Output   string `json:"output"`
}

type Version struct {
	CoreEngineVersion string `json:"core_engine_version"`
}

func GetVersion() (Version, error) {
	out, err := utils.CMD(cmd, "baseinfo")
	if err != nil {
		return Version{}, err
	}

	ver := Version{}
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		if strings.Contains(line, "Core engine version:") {
			ver.CoreEngineVersion = strings.TrimSpace(strings.TrimPrefix(line, "Core engine version:"))
			break
		}
	}

	return ver, nil
}

func ScanFile(filepath string) (Result, error) {
	out, err := utils.CMD(cmd, "scan", filepath)
	if err != nil {
		return Result{}, err
	}

	// Grab the detection result
	re := regexp.MustCompile(regexStr)
	l := re.FindStringSubmatch(out)
	res := Result{}
	if len(l) > 0 {
		res.Output = l[1]
		res.Infected = true
	}
	return res, nil
}
