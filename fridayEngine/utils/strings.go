package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

func SliceContainsString(a string, list []string) bool {
	for _, b := range list {
		if strings.Contains(b, a) {
			return true
		}
	}
	return false
}

func UniqueSlice(slice []string) []string {
	cleaned := []string{}
	for _, value := range slice {
		if !StringInSlice(value, cleaned) {
			cleaned = append(cleaned, value)
		}
	}
	return cleaned
}

func decodeUTF16(b []byte) (string, error) {

	if len(b)%2 != 0 {
		return "", fmt.Errorf("Must have even length byte slice")
	}

	u16s := make([]uint16, 1)

	ret := &bytes.Buffer{}

	b8buf := make([]byte, 4)

	lb := len(b)
	for i := 0; i < lb; i += 2 {
		u16s[0] = uint16(b[i]) + (uint16(b[i+1]) << 8)
		r := utf16.Decode(u16s)
		n := utf8.EncodeRune(b8buf, r[0])
		ret.Write(b8buf[:n])
	}

	return ret.String(), nil
}

func decode(imm uint32) string {

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, imm)
	s := fmt.Sprintf("%x", buf)
	bs, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return string(bs)
}

func check(e error) {
	if e != nil {
		logrus.Fatal(e)
	}
}

func GetASCIIStrings(data []byte, n int) []string {
	StringRegex := fmt.Sprintf("[ -~]{%d,}", n)
	re := regexp.MustCompile(StringRegex)
	asciiStrings := re.FindAllString(string(data), -1)
	return asciiStrings
}

func GetUnicodeStrings(data []byte, n int) []string {
	StringRegex := fmt.Sprintf("(?:[ -~][\x00]){%d,}", n)
	re := regexp.MustCompile(StringRegex)
	unicodeStrings := re.FindAllString(string(data), -1)

	var s []string
	for _, str := range unicodeStrings {
		decoded, _ := decodeUTF16([]byte(str))
		s = append(s, decoded)
	}
	return s
}
