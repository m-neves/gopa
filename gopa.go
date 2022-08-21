package pkg

import "fmt"

type Gopa struct{}

func New() *Gopa {
	g := new(Gopa)
	return g
}

func (*Gopa) SayHi() {
	fmt.Println("Suh")
}
