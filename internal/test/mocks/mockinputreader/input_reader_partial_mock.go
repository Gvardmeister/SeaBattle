package mockinputreader

import "errors"

type PartialErrorInputReader struct {
	calls int
}

func NewPartialErrorInputReader() *PartialErrorInputReader {
	return &PartialErrorInputReader{}
}

func (r *PartialErrorInputReader) ReadLine() (string, error) {
	if r.calls == 0 {
		r.calls++
		return "", errors.New("ошибка ввода")
	}
	return "3 4", nil
}
