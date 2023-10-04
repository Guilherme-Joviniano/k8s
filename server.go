package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/health", Health)
	http.HandleFunc("/members/admin", AdminRegisteredMembers)
	http.HandleFunc("/members", RegisteredMembers)
	http.HandleFunc("/", Greeting)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")

	message := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprint("Hello I'm the ", name, " API 😁😝"),
	}

	response, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	w.Write(response)
}

func RegisteredMembers(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("data/members.txt")

	if err != nil {
		log.Fatalf("Error reading file: %s", err.Error())
	}

	fmt.Fprintf(w, "Members: %s", string(data))
}
func AdminRegisteredMembers(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "Admin Member: USERNAME: %s \n PASSWORD: %s", user, password)
}

func Health(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
		return
	}

	res := struct {
		Message  string        `json:"message"`
		Duration time.Duration `json:"timeAlive"`
	}{
		Message:  "ok",
		Duration: duration,
	}

	response, _ := json.Marshal(res)

	w.WriteHeader(200)
	w.Write(response)
}
