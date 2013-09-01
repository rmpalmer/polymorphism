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

type Hippo struct {
	base.Pet
}

func NewHippo (name string) *Hippo {
	h := new(Hippo)
	h.SetName(name)
	return h
}

func (h *Hippo) Talk(words string) {
	fmt.Printf("Hippo named " + h.GetName() + " says " + words + "\n");
}
type Zebra struct {
	base.Pet
}

func NewZebra (name string) *Zebra {
	z := new(Zebra)
	z.SetName(name)
	return z
}

func (z *Zebra) Talk(words string) {
	fmt.Printf("Zebra named " + z.GetName() + " says " + words + "\n");
}
