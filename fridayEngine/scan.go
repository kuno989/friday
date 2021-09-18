package fridayEngine

import (
	"github.com/kuno989/friday/backend/schema"
	"github.com/kuno989/friday/fridayEngine/utils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
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
	return res
}
