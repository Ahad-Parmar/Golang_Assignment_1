package main

import (
	"encoding/json"
	"fmt"
)

type Company struct {
	C_name  string `json:"E_name"`
	C_id string `json:"E_id"`
}

type Employee struct {
	E_name      string `json:"name"`
	E_age       int    `json:"age"`
	
}

func main() {
	fmt.Println("JSON")

	employee := Employee{E_name: "Abcd", E_Age: 22}

	company := Company{C_name: "kloudone", C_id: bose}

	fmt.Printf("%+v\n", Company)

	byteArray, error := json.MarshalIndent(Company, "", "   ")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(string(byteArray))
}
