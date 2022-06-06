package main

type T struct {
	n int
}

type X struct{}

func (x X) test() {
	println(x)
}
func main() {
	var a *X
	a.test()
	X{}.test()
}
