package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	type Employee struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Address     string    `json:"address"`
		DateOfBirth time.Time `json:"dateOfBirth"`
		Job         string    `json:"job"`
		Salary      float64   `json:"salary"`
		ManagerID   string    `json:"manager-id"`
	}

	var employees = []Employee{
		{
			ID: "64133ca9-13da-4f21-ac97-5a10553272a1",
			Name: "Dilbert Maxmillion",
			Salary: 520405.50,
		},{
			ID: "64133ca9-13da-4f21-ac97-5a10553272a1",
			Name: "Sally Jackson",
			Salary: 690050.32,
			Job: "Product Architect",
		},
	}

	fmt.Println(employees[0])
	fmt.Println(employees[0].Salary)
	fmt.Println(employees[0].Name)
	fmt.Println(employees[1].Name)

	data, err := json.Marshal(employees)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", data)

}
