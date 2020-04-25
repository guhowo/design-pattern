package simple_factory

type Input struct {
	A int
	B int
}
type ICompute interface {
	Compute() int
}

