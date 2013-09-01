package base

import (

)

type Pet struct {
	name string
}
func (p *Pet) SetName (name string) {
	p.name = name
}
func (p *Pet) GetName () string {
	return p.name
}
