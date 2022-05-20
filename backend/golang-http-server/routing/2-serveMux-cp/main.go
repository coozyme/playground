package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan ServeMux

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here

	t := time.Now()
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())))

}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	var response string
	name := r.URL.Query()["name"]

	if len(name) >= 1 {
		response = `Hello, ` + name[0] + `!`
	} else {
		response = "Hello there"
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/time", TimeHandler)
	mux.HandleFunc("/time", SayHelloHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
