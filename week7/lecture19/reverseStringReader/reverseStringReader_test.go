package reverseStringReader_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"week7Lecture19Tasks/week7/lecture19/reverseStringReader"
)

func TestNewReverseStringReader(t *testing.T) {

	ourReverseStringReader := *reverseStringReader.NewReverseStringReader("ase")
	var out bytes.Buffer
	io.Copy(&out, ourReverseStringReader)
	if out.String() != "esa" {
		t.Error("Bigger card was incorrect, got" + fmt.Sprint(out) + ", want: esa")
	}
}
