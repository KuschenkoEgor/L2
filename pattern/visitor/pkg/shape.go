package pkg

type Shape interface {
	getType() string
	accept(Visitor)
}
