package stdinputreader

import (
	"bufio"
	"os"
	"strings"
)

type StdInputReader struct {
	reader *bufio.Reader
}

func NewStdInputReader() *StdInputReader {
	return &StdInputReader{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (r *StdInputReader) ReadLine() (string, error) {
	input, err := r.reader.ReadString('\n')
	return strings.TrimSpace(input), err
}
