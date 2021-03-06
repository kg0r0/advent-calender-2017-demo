package sample

import (
	"bytes"
	"strings"
	"testing"
)

func TestOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	Output(buf, "test")
	output := strings.Trim(string(buf.Bytes()), " \n\t ")
	if output != "test" {
		t.Error("unexpected value")
	}
}
