package asciiart

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func OpenFile(fn string) []byte {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func DetectFType(fn string) (d []byte, ft string) {
	b := OpenFile(fn)
	s := string(b)
	var ansiESC = regexp.MustCompile(`\[(\d)*(;)*(\d)*(;)*(\d)m`)

	if strings.HasPrefix(s, "P1") {
		return b, "PBM"
	} else if strings.HasPrefix(s, "P2") {
		return b, "PGM"
	} else if strings.HasPrefix(s, "P3") {
		return b, "PPM"
	} else if ansiESC.MatchString(s) {
		return b, "ANS"
	} else {
		return d, "ASC"
	}
}

func GetFileType(fn string) string {
	_, ft := DetectFType(fn)
	return ft
}
