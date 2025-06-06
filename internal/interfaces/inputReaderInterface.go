package interfaces

type InputReader interface {
	ReadLine() (string, error)
}
