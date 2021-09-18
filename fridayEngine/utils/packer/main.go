package packer

import (
	"github.com/kuno989/friday/fridayEngine/utils"
	"strings"
)

const tools = "/Applications/die.app/Contents/MacOS/diec"

func Scan(path string) ([]string, error) {
	args := []string{path}
	result, err := utils.CMD(tools, args...)
	if err != nil {
		return nil, err
	}
	return output(result), nil
}

func output(output string) []string {
	results := []string{}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line != "" {
			results = append(results, line)
		}
	}
	return results
}
