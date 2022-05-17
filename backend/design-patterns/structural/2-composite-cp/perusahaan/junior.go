package perusahaan

type Junior struct {
}

func (j Junior) GetSalary() int {
	return 10
}

func (j Junior) TotalDivisonSalary() int {
	// TODO: replace this
	return j.GetSalary()
}
