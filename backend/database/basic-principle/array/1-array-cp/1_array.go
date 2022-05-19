package main

type EmployeeRow struct {
	ID        int
	Name      string
	Position  string
	Salary    int
	ManagerID int
}
type EmployeeDB []EmployeeRow

func NewEmployeeDB() *EmployeeDB {
	return &EmployeeDB{}
}

func (db *EmployeeDB) Where(id int) *EmployeeRow {
	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			return &(*db)[i]
		}
	}
	return nil
}

func (db *EmployeeDB) Insert(name string, position string, salary int, managerID int) {
	(*db) = append(*db, EmployeeRow{
		ID:        len(*db) + 1,
		Name:      name,
		Position:  position,
		Salary:    salary,
		ManagerID: managerID,
	})
}

func (db *EmployeeDB) Update(id int, name string, position string, salary int, managerID int) {
	// TODO: answer here
	(*db)[id-1].ID = id
	(*db)[id-1].Name = name
	(*db)[id-1].Position = position
	(*db)[id-1].Salary = salary
	(*db)[id-1].ManagerID = managerID
}

func (db *EmployeeDB) Delete(id int) {
	// TODO: answer here

	(*db)[id-1].ID = 0
	(*db)[id-1].Name = ""
	(*db)[id-1].Position = ""
	(*db)[id-1].Salary = 0
	(*db)[id-1].ManagerID = 0

}
