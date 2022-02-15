package pkg

type IGun interface {
	SetName(Name string)
	SetPower(Power int)
	GetName() string
	GetPower() int
}
