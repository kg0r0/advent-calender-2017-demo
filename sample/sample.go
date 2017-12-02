package sample

import (
	"fmt"
	"io"
	"os"
)

func Output(w io.Writer, output string) {
	fmt.Fprintln(w, output)
}

func Sample() {
	Output(os.Stdout, "output")
}
