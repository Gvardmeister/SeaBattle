package mockinputreader

import "errors"

type MockInputReader struct {
	inputs []string
	index  int
}

func NewMockInput(inputs []string) *MockInputReader {
	return &MockInputReader{
		inputs: inputs,
	}
}

func (m *MockInputReader) ReadLine() (string, error) {
	if m.index >= len(m.inputs) {
		return "", errors.New("no more input")
	}
	input := m.inputs[m.index]
	m.index++
	return input, nil
}
