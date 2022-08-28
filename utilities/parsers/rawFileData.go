package parsers

import (
	"os"
	"strings"
)

type rawFileData struct {
	rawMeta     string
	rawMarkdown string
	fileName    string
}

func newRawFileData(fileName string) *rawFileData {
	raw, _ := os.ReadFile("./blogPosts/" + fileName + ".md")

	// Split file contents into raw parts
	dividedRaw := strings.Split(string(raw), "---")

	return &rawFileData{rawMeta: dividedRaw[1], rawMarkdown: dividedRaw[2], fileName: fileName}
}
