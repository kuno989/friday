package fridayEngine

import (
	"fmt"
	"github.com/kuno989/friday/backend/schema"
	"github.com/kuno989/friday/fridayEngine/utils"
	"github.com/kuno989/friday/fridayEngine/utils/packer"
	"github.com/saferwall/pe"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

func (s *Server) defaultScan(path string) schema.Result {
	res := schema.Result{}
	b, _ := ioutil.ReadFile(path)
	result := utils.ByteHashing(b)
	res.Size = int64(len(b))
	res.Ssdeep = result.Ssdeep
	res.Md5 = result.Md5
	res.Sha1 = result.Sha1
	res.Sha256 = result.Sha256
	res.Sha512 = result.Sha512
	res.Crc32 = result.Crc32

	logrus.Infof("file hashing finished ")
	packerRes, err := packer.Scan(path)
	if err != nil {
		logrus.Errorf("packer scan failed with %s", err)
	}

	n := 10
	asciiStrings := utils.GetASCIIStrings(b, n)
	wideStrings := utils.GetUnicodeStrings(b, n)
	asmStrings := utils.GetAsmStrings(b)
	fmt.Println(asciiStrings, wideStrings, asmStrings)

	res.Packer = packerRes
	var tags []string
	for _, out := range res.Packer {
		fmt.Println(out)
		if strings.Contains(out, "packer") ||
			strings.Contains(out, "protector") ||
			strings.Contains(out, "compiler") ||
			strings.Contains(out, "installer") ||
			strings.Contains(out, "library") {
			for sig, tag := range schema.SigMap {
				if strings.Contains(out, sig) {
					tags = append(tags, tag)
				}
			}
		}
	}
	logrus.Infof("tags extraction finish")

	//fmt.Println("tags", tags)
	//file, err := s.parser(path)
	//if err != nil {
	//	logrus.Errorf("pe parsing failed %s", err)
	//}
	//var tags []string
	//if file.IsEXE() {
	//	tags = append(tags, "exe")
	//} else if file.IsDriver() {
	//	tags = append(tags, "sys")
	//} else if file.IsDLL() {
	//	tags = append(tags, "dll")
	//}
	//res.Tags = tags

	return res
}

func (s *Server) parser(path string) (*pe.File, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	opts := pe.Options{SectionEntropy: true}
	f, err := pe.New(path, &opts)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = f.Parse()
	return f, err
}
