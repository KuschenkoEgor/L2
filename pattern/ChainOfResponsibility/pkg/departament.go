package pkg

type Departament interface {
	Execute(*Patient)
	SetNext(Departament)
}
