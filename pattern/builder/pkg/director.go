package pkg

type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{builder: b}
}

func (d *director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) BuildHouse() house {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
