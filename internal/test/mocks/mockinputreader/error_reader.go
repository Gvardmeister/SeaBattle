package mockinputreader

import "errors"

type MockErrorInputReader struct{}

func NewMockErrorInputReader() *MockErrorInputReader {
	return &MockErrorInputReader{}
}

func (m *MockErrorInputReader) ReadLine() (string, error) {
	return "", errors.New("mock read error")
}
