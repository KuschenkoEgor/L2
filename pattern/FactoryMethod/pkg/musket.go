package pkg

type Musket struct {
	Gun
}

func NewMusket() IGun{
	return &Musket{Gun{
		Name: "Musket gun",
		Power: 1,
	}}
}
