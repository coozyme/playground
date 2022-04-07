package repository

import (
	"errors"
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
	return []User{}, nil // TODO: replace this
}

func (u *UserRepository) SelectAll() ([]User, error) {
	data, _ := u.db.Load("users")

	var users []User
	for index, val := range data {
		if index != 0 {
			covertBool, _ := strconv.ParseBool(val[2])
			users = append(users, User{
				Username: val[0],
				Password: val[1],
				Loggedin: covertBool,
			})
		}
	}
	return users, nil // TODO: replace this
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	data, _ := u.db.Load("users")

	isFindUser := false
	for index, val := range data {
		if index != 0 {

			data[index][2] = "false"

			if val[0] == username && val[1] == password {
				isFindUser = true
				data[index][2] = "true"
			}
		}
	}

	if isFindUser == false {
		return nil, errors.New("Login Failed")
	}

	err := u.db.Save("users", data)
	if err != nil {
		return nil, err
	}

	return &username, nil
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	data, _ := u.db.Load("users")

	isFindUserLogin := false
	username := ""
	for index, val := range data {
		if index != 0 {
			covertBool, err := strconv.ParseBool(val[2])
			if err != nil {
				panic("Errror")
			}
			if covertBool == true {
				isFindUserLogin = true
				username = val[0]
				break
			}
		}
	}

	if isFindUserLogin == false {
		return nil, errors.New("no user is logged in")
	}
	return &username, nil
}

func (u *UserRepository) Logout(username string) error {
	data, _ := u.db.Load("users")

	isFindUserLogin := false

	for index, val := range data {
		if index != 0 {
			if username == val[0] {
				if val[2] == "false" {
					break
				} else {
					data[index][2] = "false"
					isFindUserLogin = true
					break
				}
			}
		}
	}
	if isFindUserLogin == false {
		return errors.New("no user is logged in")
	}

	err := u.db.Save("users", data)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Save(users []User) error {
	data, _ := u.db.Load("users")

	for _, val := range users {
		newUser := []string{val.Username, val.Password, strconv.FormatBool(val.Loggedin)}
		data = append(data, newUser)
	}

	err := u.db.Save("users", data)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	data, _ := u.db.Load("users")

	isFindUser := false
	for index, val := range data {
		if val[0] == username {
			isFindUser = true
			data[index][2] = strconv.FormatBool(status)
			break
		}
	}

	if isFindUser == false {
		return errors.New("Login Failed")
	}
	return nil
}

func (u *UserRepository) LogoutAll() error {
	data, _ := u.db.Load("users")

	for index, _ := range data {
		if index != 0 {
			data[index][2] = "false"
		}
	}

	u.db.Save("users", data)

	return nil
}
