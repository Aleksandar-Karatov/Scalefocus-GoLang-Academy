package reverseStringReader

import (
	"io"
)

type ReverseStringReader interface {
	Read(p []byte) (int, error)
}

type OurReverseStringReader struct {
	index    int64
	prevRune rune
	buff     []byte
}

func NewReverseStringReader(input string) *ReverseStringReader {
	var ourRevReader OurReverseStringReader
	ourRevReader.buff = []byte(input)
	for i, j := 0, len(ourRevReader.buff)-1; i < j; i, j = i+1, j-1 {
		ourRevReader.buff[i], ourRevReader.buff[j] = ourRevReader.buff[j], ourRevReader.buff[i]
	}
	rvReader := ReverseStringReader(&ourRevReader)
	return &rvReader
}

func (rd *OurReverseStringReader) Read(p []byte) (n int, err error) {
	if rd.index >= int64(len(rd.buff)) {
		return 0, io.EOF
	}
	rd.prevRune = -1
	n = copy(p, rd.buff[rd.index:])
	rd.index += int64(n)
	return

}
