package main

import (
	"fmt"
	"time"
)

func main() {

	type Employee struct {
		ID          string
		Name        string
		Address     string
		DateOfBirth time.Time
		Job         string
		Salary      float64
		ManagerID   string
	}

	var dilbert Employee

	dilbert.Salary = 50602.66
	dilbert.Name = "Dilbert Maxmilliarn"
	dilbert.Job = "Engineer Genius"
	dilbert.Address = "That place in the cartoon."

	fmt.Println(dilbert)
	fmt.Println(dilbert.Salary)
	fmt.Println(dilbert.Name)
}
