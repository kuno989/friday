package schema

import (
	"github.com/saferwall/pe"
	"time"
)

type FileResponse struct {
	Sha256      string `json:"sha256,omitempty"`
	FileName    string `json:"filename,omitempty"`
	FileSize    int64  `json:"filesize,omitempty"`
	Message     string `json:"message,omitempty"`
	Description string `json:"description,omitempty"`
}

type Submission struct {
	Date     *time.Time `json:"date,omitempty"`
	Filename string     `json:"filename,omitempty"`
	Country  string     `json:"country,omitempty"`
}

type File struct {
	FileKey          string                 `json:"file_key,omitempty"`
	Md5              string                 `json:"md5,omitempty"`
	Sha1             string                 `json:"sha1,omitempty"`
	Sha256           string                 `json:"sha256,omitempty"`
	Sha512           string                 `json:"sha512,omitempty"`
	Ssdeep           string                 `json:"ssdeep,omitempty"`
	Crc32            string                 `json:"crc32,omitempty"`
	Magic            string                 `json:"magic,omitempty"`
	Size             int64                  `json:"size,omitempty"`
	Exif             map[string]string      `json:"exif,omitempty"`
	Tags             map[string]interface{} `json:"tags,omitempty"`
	Packer           []string               `json:"packer,omitempty"`
	FirstSubmission  *time.Time             `json:"first_submission,omitempty"`
	LastSubmission   *time.Time             `json:"last_submission,omitempty"`
	LastScanned      *time.Time             `json:"last_scanned,omitempty"`
	Submissions      []Submission           `json:"submissions,omitempty"`
	SubmissionsCount int64                  `json:"submissions_count"`
	Strings          []StringStruct         `json:"strings,omitempty"`
	MultiAV          map[string]interface{} `json:"multiav,omitempty"`
	Status           int                    `json:"status,omitempty"`
	PE               *pe.File               `json:"pe,omitempty"`
	Histogram        []int                  `json:"histogram,omitempty"`
	ByteEntropy      []int                  `json:"byte_entropy,omitempty"`
	Type             string                 `json:"type,omitempty"`
	// Comments        []Comment              `json:"comments,omitempty"`
}

type StringStruct struct {
	Encoding string `json:"encoding"`
	Value    string `json:"value"`
}

type Result struct {
	Md5         string                 `json:"md5,omitempty"`
	Sha1        string                 `json:"sha1,omitempty"`
	Sha256      string                 `json:"sha256,omitempty"`
	Sha512      string                 `json:"sha512,omitempty"`
	Ssdeep      string                 `json:"ssdeep,omitempty"`
	Crc32       string                 `json:"crc32,omitempty"`
	Magic       string                 `json:"magic,omitempty"`
	Size        int64                  `json:"size,omitempty"`
	Exif        map[string]string      `json:"exif,omitempty"`
	Tags        map[string]interface{} `json:"tags,omitempty"`
	Packer      []string               `json:"packer,omitempty"`
	LastScanned *time.Time             `json:"last_scanned,omitempty"`
	Strings     []StringStruct         `json:"strings,omitempty"`
	MultiAV     map[string]interface{} `json:"multiav,omitempty"`
	Status      int                    `json:"status,omitempty"`
	PE          *pe.File               `json:"pe,omitempty"`
	Histogram   []int                  `json:"histogram,omitempty"`
	ByteEntropy []int                  `json:"byte_entropy,omitempty"`
	Type        string                 `json:"type,omitempty"`
}

var SigMap = map[string]string{
	"Nullsoft Scriptable Install System": "nsis",
	"Inno Setup":                         "innosetup",
	"UPX":                                "upx",
	"FSG":                                "fsg",
	"ASPack":                             "aspack",
	"RLPack":                             "rlpack",
	"ASProtect":                          "asprotect",
	"ACProtect":                          "acprotect",
	"PECompact":                          "pecompact",
	"PECrypt32":                          "pecrypt32",
	"PE-Armor":                           "pe-armor",
	"Petite":                             "petite",
	"PELock":                             "pelock",
	"tElock":                             "telock",
	"EXECryptor":                         "execryptor",
	"Obsidium":                           "obsidium",
	"VMProtect":                          "vmprotect",
	"Themida/Winlicense":                 "themida-winlicense",
	"MoleBox":                            "molebox",
	"ENIGMA":                             "enigma",
	"MPRESS":                             "mpress",
	"Armadillo":                          "armadillo",
	"gcc":                                "gcc",
	"MinGW":                              "mingw",
	"Microsoft Visual C/C++":             "vc",
	"Microsoft Visual Basic":             "vb",
	"Borland C++":                        "borland-c++",
	"MASM":                               "masm",
	"FASM":                               "fasm",
	".NET":                               "dotnet",
	"MFC":                                "mfc",
	"Yoda's Crypter":                     "yodascrypter",
	"Delphi":                             "delphi",
	"AutoIt":                             "autoit",
	"StarForce":                          "starforce",
	"eXPressor":                          "expressor",
	"sfx: Microsoft Cabinet":             "sfx-cab",
	"Smart Assembly":                     "smart-assembly",
	".NET Reactor":                       "dotnet-reactor",
	"Confuser":                           "confuser",
	"Dotfuscator":                        "dotfuscator",
}
