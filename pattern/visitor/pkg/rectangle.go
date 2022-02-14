package pkg

type Rectangle struct {
	L int
	B int
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}
