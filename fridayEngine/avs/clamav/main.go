package clamav

import (
	"github.com/kuno989/friday/fridayEngine/utils"
	"strings"
)

const (
	clamavscan = "/usr/bin/clamdscan"
)

type Result struct {
	Infected bool   `json:"infected"`
	Output   string `json:"output"`
}

func ScanFile(path string) (Result, error) {
	out, err := utils.CMD(clamavscan, "--no-summary", path)
	if err != nil && err.Error() != "exit status 1" {
		return Result{}, err
	}

	if strings.HasSuffix(out, "OK\n") {
		return Result{}, nil
	}
	if !strings.HasSuffix(out, "FOUND\n") {
		return Result{}, nil
	}

	res := Result{}
	parts := strings.Split(out, ": ")
	det := parts[len(parts)-1]
	res.Output = strings.TrimSuffix(det, " FOUND\n")
	res.Infected = true
	return res, nil
}

func GetVersion() (string, error) {
	out, err := utils.CMD(clamavscan, "--version")
	if err != nil {
		return "", err
	}
	ver := strings.Split(out, "/")[0]
	ver = strings.Split(ver, " ")[1]
	return ver, nil
}
