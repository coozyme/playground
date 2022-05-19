package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Pokemon struct {
	Name    string `json:"name"`
	Species struct {
		Name string `json:"name"`
	} `json:"species"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		}
	}
}

func GetPokemonData() (*Pokemon, error) {
	apiPath := "https://pokeapi.co/api/v2/pokemon/1"
	// TODO: answer here

	var resp Pokemon

	response, err := http.Get(apiPath)

	if err != nil {
		log.Fatalf("failed get api, err : %v", err)
		return &resp, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed ReadAll, err: %v", err)
		return &resp, err
	}

	json.Unmarshal(responseData, &resp)

	return &resp, nil

	panic("Not yet implemented")
}

func main() {
	pokemon, err := GetPokemonData()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pokemon)
}
