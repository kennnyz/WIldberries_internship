package main

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setNumFloor()
	d.builder.setWindowType()
	d.builder.setDoorType()
	return d.builder.getHouse()
}
