package repository

import (
	"fmt"
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	record, err := u.db.Load("users")
	if err != nil {
		record = [][]string{
			{"username", "password", "loggedin"},
		}
		if err := u.db.Save("users", record); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(record); i++ {
		loggedin, err := strconv.ParseBool(record[i][2])
		if err != nil {
			return nil, err
		}
		user := User{
			Username: record[i][0],
			Password: record[i][1],
			Loggedin: loggedin,
		}
		result = append(result, user)
	}
	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {

	if err := u.LogoutAll(); err != nil {
		return nil, err
	}

	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username && user.Password == password {
			u.changeStatus(username, true)
			return &user.Username, nil
		}
	}
	return nil, fmt.Errorf("Login Failed")
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Loggedin {
			return &user.Username, nil
		}
	}
	return nil, fmt.Errorf("no user is logged in")
}

func (u *UserRepository) Logout(username string) error {
	userLogin, err := u.FindLoggedinUser()
	if err != nil {
		return err
	}
	if userLogin == nil {
		return fmt.Errorf("no user is logged in")
	}
	return u.changeStatus(*userLogin, false)
}

func (u *UserRepository) Save(users []User) error {
	records := [][]string{
		{"username", "password", "loggedin"},
	}
	for i := 0; i < len(users); i++ {
		records = append(records, []string{
			users[i].Username,
			users[i].Password,
			strconv.FormatBool(users[i].Loggedin),
		})
	}
	return u.db.Save("users", records)
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i].Loggedin = status
			return u.Save(users)
		}
	}
	return fmt.Errorf("User not found")
}

func (u *UserRepository) LogoutAll() error {
	users, err := u.SelectAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		u.Logout(user.Username)
	}
	return nil
}
