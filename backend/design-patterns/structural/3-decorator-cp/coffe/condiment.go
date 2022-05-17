package coffe

type Mocha struct {
	// TODO: answer here
	Coffe Coffe
}

func (m Mocha) GetCost() float64 {
	// TODO: replace this
	return m.Coffe.GetCost() + 1.00
}

func (m Mocha) GetDescription() string {
	// TODO: replace this
	return m.Coffe.GetDescription() + ", Mocha"
}

type Whipcream struct {
	// TODO: answer here
	Coffe Coffe
}

func (w Whipcream) GetCost() float64 {
	// TODO: replace this
	return w.Coffe.GetCost() + 0.20
}

func (w Whipcream) GetDescription() string {
	// TODO: replace this
	return w.Coffe.GetDescription() + ", Whipcream"
}
