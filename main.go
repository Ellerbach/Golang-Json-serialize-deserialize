package main

import (
	"encoding/json"
	"fmt"

	null "gopkg.in/guregu/null.v3"
)

// Employee class example
type Employee struct {
	Name   null.String `json:"name,omitempty"`
	Age    null.Int    `json:"age"`
	Salary int         `json:"salary"`
	// In GO, no recursivity is permitted,
	// So need to use pointers to avoid it
	Employee *Employee `json:"employee"`
	Other    *Other    `json:"other"`
}

// Employee2 class example
type Employee2 struct {
	Name     string     `json:"name,omitempty"`
	Age      int        `json:"age,omitempty"`
	Salary   int        `json:"salary,omitempty"`
	Employee *Employee2 `json:"employee,omitempty"`
	Other    *Other     `json:"other,omitempty"`
}

// Employee3 class example
type Employee3 struct {
	Name     string     `json:"name"`
	Age      int        `json:"age"`
	Salary   int        `json:"salary"`
	Employee *Employee3 `json:"employee"`
	Other    *Other     `json:"other"`
}

// Employee4 class example
type Employee4 struct {
	Name     string     `json:"name,omitempty"`
	Age      int        `json:"age,omitempty"`
	Salary   int        `json:"salary,omitempty"`
	Employee *Employee4 `json:"employee,omitempty"`
	Other    Other      `json:"other,omitempty"`
}

// Other structure
type Other struct {
	Tagada string `json:"tagada,omitempty"`
}

func main() {
	fmt.Println("Json serialisation and deserialization in GO")
	empobj := Employee{Age: null.IntFrom(24), Salary: 0}
	empobj2 := Employee2{Name: "", Age: 24, Salary: 0}
	empobj3 := Employee3{Name: "", Age: 24, Salary: 0}
	empobj4 := Employee4{Name: "", Age: 24, Salary: 0}
	emp, _ := json.Marshal(empobj)
	emp2, _ := json.Marshal(empobj2)
	emp3, _ := json.Marshal(empobj3)
	emp4, _ := json.Marshal(empobj4)
	fmt.Println(string(emp))
	fmt.Println(string(emp2))
	fmt.Println(string(emp3))
	fmt.Println(string(emp4))

	dt := "{\"name\": null, \"age\": null, \"salary\": 0, \"employee\": { \"name\": \"tagada\", \"age\": 25, \"salary\": 34567}}"
	emptd := []byte(dt)
	fmt.Println(string(emptd))
	var retobj Employee
	var retobj2 Employee2
	var retobj3 Employee3
	var retobj4 Employee4
	json.Unmarshal(emptd, &retobj)
	json.Unmarshal(emptd, &retobj2)
	json.Unmarshal(emptd, &retobj3)
	json.Unmarshal(emptd, &retobj4)
	fmt.Println("From Employee, Age, Name, Employee and Other")
	fmt.Println(retobj.Age)
	fmt.Println(retobj.Name)
	fmt.Println(retobj.Employee)
	fmt.Println(retobj.Other)
	fmt.Println("From Employee2, Age, Name, Employee and Other")
	fmt.Println(retobj2.Age)
	fmt.Println(retobj2.Name)
	fmt.Println(retobj2.Employee)
	fmt.Println(retobj2.Other)
	fmt.Println("From Employee3, Age, Name, Employee and Other")
	fmt.Println(retobj3.Age)
	fmt.Println(retobj3.Name)
	fmt.Println(retobj3.Employee)
	fmt.Println(retobj3.Other)
	fmt.Println("From Employee4, Age, Name, Employee and Other")
	fmt.Println(retobj4.Age)
	fmt.Println(retobj4.Name)
	fmt.Println(retobj4.Employee)
	fmt.Println(retobj4.Other)
	fmt.Println("From Employee, reserializing it")
	ser, _ := json.Marshal(retobj)
	fmt.Println(string(ser))
	fmt.Println("From Employee2, reserializing it")
	ser2, _ := json.Marshal(retobj2)
	fmt.Println(string(ser2))
	fmt.Println("From Employee3, reserializing it")
	ser3, _ := json.Marshal(retobj3)
	fmt.Println(string(ser3))
	fmt.Println("From Employee4, reserializing it")
	ser4, _ := json.Marshal(retobj4)
	fmt.Println(string(ser4))
}
