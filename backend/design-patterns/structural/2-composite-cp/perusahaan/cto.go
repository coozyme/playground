package perusahaan

type CTO struct {
	Subordinate []Employee
}

func (c CTO) GetSalary() int {
	return 30
}

func (c CTO) TotalDivisonSalary() int {
	// TODO: replace this
	totalSalary := 0

	for _, data := range c.Subordinate {
		totalSalary += data.TotalDivisonSalary()
	}
	return totalSalary + c.GetSalary()
}
