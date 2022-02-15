package pkg



type Gun struct {
	Name string
	Power int
}

func (g *Gun) SetName(Name string){
	g.Name = Name
}

func (g *Gun) GetName() string{
	return g.Name
}

func (g *Gun) SetPower(Power int){
	g.Power = Power
}

func (g *Gun) GetPower() int{
	return g.Power
}