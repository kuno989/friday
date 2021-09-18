package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/bnagy/gapstone"
	"github.com/sirupsen/logrus"
	"regexp"
	"unicode/utf16"
	"unicode/utf8"
)

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

// GetUnicodeStrings returns list of Unicode strings
// n: defines minimum length of string
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

// GetAsmStrings returns list of stacked strings
// Well this is not finished, need a lot of enhancements.
func GetAsmStrings(x86Code32 []byte) (result []string) {

	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("Asm string extraction failed: %v\n", err)
		}
	}()

	engine, err := gapstone.New(
		gapstone.CS_ARCH_X86,
		gapstone.CS_MODE_32,
	)
	check(err)

	defer engine.Close()

	// maj, min := engine.Version()
	// log.Printf("Hello Capstone! Version: %v.%v\n", maj, min)

	err = engine.SetOption(gapstone.CS_OPT_DETAIL, gapstone.CS_OPT_ON)
	check(err)

	for offset := 0; offset < len(x86Code32); offset++ {
		var buffer bytes.Buffer
		var countConcat = 0
		if x86Code32[offset] == 0xC7 && (x86Code32[offset+1] == 0x45 || x86Code32[offset+1] == 0x84 || x86Code32[offset+1] == 0x85 || x86Code32[offset+1] == 0x44) {

			// log.Printf("Found a 0xC7 at offset: 0x%x", offset)
			//disassemble arround it
			insns, err := engine.Disasm(x86Code32[offset:], uint64(offset), 20)
			check(err)

			// iterater over instructions
			for i, insn := range insns {

				// log.Printf("Current offset %x", offset)
				// log.Printf("0x%x:\t%s\t\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
				// let's see if the disassembled instructions looks similar
				if insns[i].Bytes[0] == insns[0].Bytes[0] && insns[i].Bytes[1] == insns[0].Bytes[1] {

					s := decode(uint32(insns[i].X86.Operands[1].Imm))
					buffer.WriteString(s)

					// Increment the offset
					offset = offset + int(insn.Size)

					countConcat = countConcat + 1
				} else {
					break
				}
			}

			// check how many concats do we have
			if countConcat > 1 {
				s := bytes.Trim(buffer.Bytes(), "\x00")
				result = append(result, string(s))
				// fmt.Printf("\n%s, %x", string(s), offset)
			}

			offset = offset - 1
		}
	}

	// elapsed := time.Since(start)
	// log.Printf("Execution took %s", elapsed)

	return result
}
