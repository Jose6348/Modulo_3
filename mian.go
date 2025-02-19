package main

type Foo struct {
	Bar string
	Baz float64
	quz []string
}

func main() {
	f := Foo{
		Bar: "",
		Baz: 0,
		quz: []string{},
	}
}
