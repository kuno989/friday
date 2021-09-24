package utils

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func RunCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		_, _ = io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

func CMD(app string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	cmd := exec.CommandContext(ctx, app, args...)

	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("시간 초과")
		return string(output), err
	}
	return string(output), err
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func ReadAll(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
