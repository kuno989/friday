package fridayEngine

import (
	"github.com/kuno989/friday/backend/schema"
	"github.com/kuno989/friday/fridayEngine/utils"
	"github.com/kuno989/friday/fridayEngine/utils/exif"
	"github.com/kuno989/friday/fridayEngine/utils/magic"
	"github.com/kuno989/friday/fridayEngine/utils/packer"
	"github.com/saferwall/pe"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

func (s *Server) defaultScan(path string, res *schema.Result) {
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
	magic, err := magic.Scan(path)
	if err != nil {
		logrus.Errorf("magic scan failed with %s", err)
	}
	res.Magic = magic
	packerRes, err := packer.Scan(path)
	if err != nil {
		logrus.Errorf("packer scan failed with %s", err)
	}
	exif, err := exif.Scan(path)
	if err != nil {
		logrus.Errorf("exif scan failed with %s", err)
	}
	res.Exif = exif

	// Extract strings.
	n := 10
	asciiStrings := utils.GetASCIIStrings(b, n)
	wideStrings := utils.GetUnicodeStrings(b, n)
	asmStrings := utils.GetAsmStrings(b)
	// 중복 제거
	uniqueASCII := utils.UniqueSlice(asciiStrings)
	uniqueWide := utils.UniqueSlice(wideStrings)
	uniqueAsm := utils.UniqueSlice(asmStrings)
	var strResults []schema.StringStruct
	for _, str := range uniqueASCII {
		strResults = append(strResults, schema.StringStruct{"ascii", str})
	}

	for _, str := range uniqueWide {
		strResults = append(strResults, schema.StringStruct{"wide", str})
	}

	for _, str := range uniqueAsm {
		strResults = append(strResults, schema.StringStruct{"asm", str})
	}
	res.Strings = strResults
	logrus.Infof("strings extraction finish")
	res.Packer = packerRes
	var tags []string
	for _, out := range res.Packer {
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

	s.parseFile(b, path, res)
	file, err := s.parser(path)
	if err != nil {
		logrus.Errorf("pe parsing failed %s", err)
	}
	s.getTags(file, res)
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

func (s *Server) getTags(file *pe.File, res *schema.Result) {
	var tags []string
	res.Tags = map[string]interface{}{}
	if file.IsEXE() {
		tags = append(tags, "exe")
	} else if file.IsDriver() {
		tags = append(tags, "sys")
	} else if file.IsDLL() {
		tags = append(tags, "dll")
	}
	if tags != nil {
		res.Tags[res.Type] = tags
	}
}

func (s *Server) parseFile(b []byte, filePath string, res *schema.Result) {
	magic := res.Magic
	if strings.HasPrefix(magic, "PE32") {
		res.Type = "pe"
	} else if strings.HasPrefix(magic, "MS-DOS") {
		res.Type = "msdos"
	} else if strings.HasPrefix(magic, "XML") {
		res.Type = "xml"
	} else if strings.HasPrefix(magic, "HTML") {
		res.Type = "html"
	} else if strings.HasPrefix(magic, "ELF") {
		res.Type = "elf"
	} else if strings.HasPrefix(magic, "PDF") {
		res.Type = "pdf"
	} else if strings.HasPrefix(magic, "Macromedia Flash") {
		res.Type = "swf"
	} else if strings.HasPrefix(magic, "Zip archive data") {
		res.Type = "zip"
	} else if strings.HasPrefix(magic, "Java archive data (JAR)") {
		res.Type = "jar"
	} else if strings.HasPrefix(magic, "JPEG image data") {
		res.Type = "jpeg"
	} else if strings.HasPrefix(magic, "PNG image data") {
		res.Type = "png"
	} else if strings.HasPrefix(magic, "SVG Scalable Vector") {
		res.Type = "svg"
	}

	var err error
	switch res.Type {
	case "pe":
		res.PE, err = parsePE(filePath)
		if err != nil {
			logrus.Errorf("pe parser failed: %v", err)
		}
		//res.Histogram = bs.ByteHistogram(b)
		//res.ByteEntropy = bs.ByteEntropyHistogram(b)
		//logrus.Debug("bytestats pkg success")
	}
}

func parsePE(filePath string) (*pe.File, error) {
	var err error
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	opts := pe.Options{SectionEntropy: true}
	pe, err := pe.New(filePath, &opts)
	if err != nil {
		return nil, err
	}
	defer pe.Close()
	// Parse the PE.
	err = pe.Parse()
	logrus.Debug("pe pkg success")
	return pe, err
}
