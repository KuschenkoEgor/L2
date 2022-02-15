package pkg

type Ak47 struct {
	Gun
}

func NewAk47() IGun{
	return &Ak47{Gun{
		Name: "AK 47 gun",
		Power: 4,
	}}
}
