package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"log"

	"github.com/glaslos/ssdeep"
)

type CryptResult struct {
	Md5    string
	Ssdeep string
	Sha1   string
	Sha256 string
	Sha512 string
	Crc32  string
}

func getMd5(b []byte) string {
	hash := md5.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}

func getSha1(b []byte) string {
	hash := sha1.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}

func getSha256(b []byte) string {
	hash := sha256.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}

func getSha512(b []byte) string {
	hash := sha512.New()
	hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}

func getCrc32(b []byte) string {
	checkSum := crc32.ChecksumIEEE(b)
	hash := fmt.Sprintf("0x%v", checkSum)
	return hash
}

func getSsdeep(b []byte) (string, error) {
	return ssdeep.FuzzyBytes(b)
}

func ByteHashing(b []byte) CryptResult {
	fuzzy, err := getSsdeep(b)
	if err != nil && err != ssdeep.ErrFileTooSmall {
		log.Fatalf("ssdeep 실패 %s", err)
	}
	result := CryptResult{
		Md5:    getMd5(b),
		Ssdeep: fuzzy,
		Sha1:   getSha1(b),
		Sha256: getSha256(b),
		Sha512: getSha512(b),
		Crc32:  getCrc32(b),
	}
	return result
}
