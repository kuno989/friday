package magic

import (
	"github.com/kuno989/friday/fridayEngine/utils"
	"strings"
)

const tools = "file"

func Scan(path string) (string, error) {
	args := []string{path}
	result, err := utils.CMD(tools, args...)
	if err != nil {
		return "", err
	}
	return output(result), nil
}

func output(output string) string {
	lines := strings.Split(output, ": ")
	if len(lines) > 0 {
		return strings.TrimSuffix(lines[1], "\n")
	}
	return ""
}
