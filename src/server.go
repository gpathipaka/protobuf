package main

import (
	fmt "fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func main() {
	log.Println("Main() start...")
	http.HandleFunc("/", hanler)
	http.ListenAndServe(":4001", nil)
	log.Println("Main() end...")
	fmt.Println()
}

func hanler(w http.ResponseWriter, r *http.Request) {
	log.Println("request Received.... ")
	ba, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	emp := Employee{}
	if err = proto.Unmarshal(ba, &emp); err != nil {
		log.Fatal(err)
	}
	log.Println(emp.FirstName, ":", emp.LastName, ":", emp.EmployeeId, ":", emp.Age)
}
