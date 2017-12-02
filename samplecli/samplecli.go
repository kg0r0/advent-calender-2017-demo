package samplecli

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Sample interface {
	Output(w io.Writer, output string)
}

func (s *Sample) Output(w io.Writer, output string) {
	fmt.Fprintln(w, str)
}

func SampleCli(args []string) {
	if len(args) < 1 {
		err := errors.New("error")
	}
	if err != nil {
		Output(os.Stderr, err.Error())
		return
	}
	for _, input := range args {
		Output(os.Stdout, input)
	}
}
