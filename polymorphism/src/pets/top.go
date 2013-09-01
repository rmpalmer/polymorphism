// this is a project that I am experimenting with
// again

package main

import (
	"fmt"
	"base"
	"zoo"
)

type Communicate interface {
	Talk(words string)
}
type Cat struct {
	base.Pet
}
func NewCat (name string) *Cat {
	c := new(Cat)
	c.SetName(name)
	return c
}
func (c *Cat) Talk(words string) {
	fmt.Printf("Cat named " + c.GetName() + " says " + words + "\n");
}
type Dog struct {
	base.Pet
}
func NewDog (name string) *Dog {
	d := new(Dog)
	d.SetName(name)
	return d
}
func (d *Dog) Talk(words string) {
	fmt.Printf("Dog named " + d.GetName() + " says " + words + "\n");
}

func DoTalk(x Communicate, words string) {
	x.Talk(words)
}
func main() {
	var c, d, g Communicate
	c = NewCat("KC")
	d = NewDog("Red")
	g = zoo.NewGiraffe("Stretch")
	c.Talk("meow MEOW")
	d.Talk("woof WOOF")
	g.Talk("nothing at all")
}
