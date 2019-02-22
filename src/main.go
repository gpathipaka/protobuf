package main

import (
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	log.Println("Main() start")
	emp := &Employee{
		EmployeeId: 1001,
		FirstName:  "Ishu",
		LastName:   "Pathipaka",
		Age:        27,
	}

	//Marshaling Employee
	empByteArr, err := proto.Marshal(emp)
	if err != nil {
		log.Fatal("Error Marshling Employee ", err)
	}
	log.Println(empByteArr)

	//Unmarshaling Employee
	emp1 := &Employee{}
	err = proto.Unmarshal(empByteArr, emp1)
	if err != nil {
		log.Fatal("Error UnMarshaling emp data..", err)
	}
	log.Println(emp1)
	log.Println("Main() end..")
}
