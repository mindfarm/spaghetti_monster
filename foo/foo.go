package foo

type Foo struct{}

func (f *Foo) Adder(a, b int) int {
	return a + b
}
