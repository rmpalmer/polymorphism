package zoo

import (
	"fmt"
	"base"
)

type Giraffe struct {
	base.Pet
}

func NewGiraffe (name string) *Giraffe {
	g := new(Giraffe)
	g.SetName(name)
	return g
}
func (g *Giraffe) Talk(words string) {
	fmt.Printf("Giraffe named " + g.GetName() + " says " + words + "\n");
}
