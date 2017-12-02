package samplecli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

const testOutput = "success"

var s Samplecli

type sampleMock struct {
}

func (sm *sampleMock) Output(w io.Writer, output string) {
	ftm.Fprintln(os.Stdout, testOutput)
}

func TestOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	s.Output(buf, testOutput)
	require.Equal(buf, []byte(testOutput))
}

func ExampleSampleCli() {
	normalArgs := []string{"", ""}
	errorArgs := []string{}
	SampleCli(testArgs)
	// Output: success
	SampleCli(errorArgs)
	// Output: error
}
