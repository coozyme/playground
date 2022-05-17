package perusahaan

type VP struct {
	Subordinate []Employee
}

func (vp VP) GetSalary() int {
	return 20
}

func (vp VP) TotalDivisonSalary() int {
	// TODO: replace this
	totalSalary := 0

	for _, data := range vp.Subordinate {
		totalSalary += data.GetSalary()
	}
	return totalSalary + vp.GetSalary()
}
