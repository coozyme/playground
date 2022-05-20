package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {
	// TODO: replace this
	result := make(map[int]int)

	for _, v := range people {
		if result[v.age] == 0 {
			result[v.age] = 1
		} else {
			result[v.age] = result[v.age] + 1
		}
	}

	fmt.Printf("ruslt %v", result)
	return result
}

func FilterByAge(people []Person, age int) []Person {
	// TODO: replace this
	result := []Person{}
	for _, v := range people {
		if v.age == age {
			result = append(result, v)
		}
	}
	return result
}
