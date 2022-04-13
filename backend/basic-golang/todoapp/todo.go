package main

import (
	"fmt"
	"time"
)

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	//beginanswer
	todos.items = append(todos.items, item)
	//endanswer
}

func (todos *Todos) GetAll() []Item {
	//beginanswer
	return todos.items
	//endanswer return []Item{}
}

func (todos *Todos) GetUpcoming() []Item {
	//beginanswer
	var upcoming []Item
	for _, item := range todos.items {
		if item.Deadline.After(time.Now()) {
			upcoming = append(upcoming, item)
		}
	}
	return upcoming
	//endanswer return []Item{}
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}

func main() {
	myTodo := NewTodos()
	myTodo.Add(NewItem("buy milk", time.Now().Add(time.Hour)))
	myTodo.Add(NewItem("buy bag", time.Now().Add(time.Hour)))

	getAll := myTodo.GetAll()

	fmt.Println(getAll)
}
