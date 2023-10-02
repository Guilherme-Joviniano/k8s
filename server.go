package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Greeting)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")

	message := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprint("Hello I'm the %s API 😁😝", name),
	}

	response, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	w.Write(response)
}
