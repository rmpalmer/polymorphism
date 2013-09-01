// this is a project that I am experimenting with

package main

import "fmt"

type Communicate interface {
	Talk(words string)
}
type Pet struct {
	name string
}
func (p *Pet) SetName (name string) {
	p.name = name
}
func (p *Pet) GetName () string {
	return p.name
}
type Cat struct {
	Pet
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
	Pet
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
	var c, d Communicate
	c = NewCat("KC")
	d = NewDog("Red")
	c.Talk("meow MEOW")
	d.Talk("woof WOOF")
}
