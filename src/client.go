package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func main() {
	log.Println("Main() start...")
	emp := &Employee{
		EmployeeId: 007,
		FirstName:  "James",
		LastName:   "Bond",
		Age:        35,
	}

	ba, err := proto.Marshal(emp)
	if err != nil {
		log.Fatal(err)
	}

	resp, httpErr := http.Post("http://localhost:4001", "", bytes.NewBuffer(ba))
	if httpErr != nil {
		log.Fatal(httpErr)
		return
	}
	log.Println("Reponse Received from Server...", resp.Body)
	log.Println("Main() end...")
}
