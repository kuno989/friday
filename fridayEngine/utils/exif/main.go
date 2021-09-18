package exif

import (
	"github.com/kuno989/friday/fridayEngine/utils"
	"strings"

	str "github.com/stoewer/go-strcase"
)

const tools = "exiftool"

func Scan(path string) (map[string]string, error) {
	args := []string{path}
	result, err := utils.CMD(tools, args...)
	if err != nil {
		return nil, err
	}
	return output(result), nil
}

func output(output string) map[string]string {
	var ignoreTags = []string{
		"Directory",
		"File Name",
		"File Permissions",
	}

	lines := strings.Split(output, "\n")

	if utils.StringInSlice("File not found", lines) {
		return nil
	}

	data := make(map[string]string, len(lines))

	for _, line := range lines {
		keyvalue := strings.Split(line, ":")
		if len(keyvalue) != 2 {
			continue
		}
		if !utils.StringInSlice(strings.TrimSpace(keyvalue[0]), ignoreTags) {
			data[strings.TrimSpace(str.UpperCamelCase(keyvalue[0]))] = strings.TrimSpace(keyvalue[1])
		}
	}

	return data
}
