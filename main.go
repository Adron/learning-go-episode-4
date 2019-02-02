package main

import "fmt"

func main() {
	ages := map[string]int{
		"Peterson": 52,
		"Sally":    22,
		"Javovia":  15,
		"Ben":      42,
	}

	jobAssociation := make(map[string]string)
	jobAssociation["Peterson"] = "Engineer"
	jobAssociation["Sally"] = "CEO"
	jobAssociation["Jovovia"] = "Gamer"
	jobAssociation["Ben"] = "Programmer"

	printAges(ages)
	printJobAssociations(jobAssociation)

	fmt.Println(ages)
	fmt.Println(jobAssociation)
	fmt.Println(jobAssociation["Jovovia"])
	fmt.Println(jobAssociation["Frank"]) // Blank! :o
	fmt.Println(ages["Sally"])

	delete(ages, "Sally")

	fmt.Println(ages)
	fmt.Println(ages["Sally"])

	delete(jobAssociation, "Jovovia")

	fmt.Println(jobAssociation)
	fmt.Println(jobAssociation["Jovovia"]) // Blank.

	ages2 := map[string]int{
		"Frank":     52,
		"Johnson":   22,
		"Smith":     15,
		"Jezebelle": 42,
	}

	ages3 := map[string]int{
		"Frank":     52,
		"Johnson":   22,
		"Smith":     15,
		"Jezebelle": 42,
	}

	if equal(ages, ages2) {
		fmt.Println("Naw, not really equal.")
	} else {
		fmt.Println("This is correct, not equal.")
	}

	if equal(ages2, ages3) {
		fmt.Println("True, these are effectively the same map values and keys.")
	}
}

func printJobAssociations(associations map[string]string) {
	for name, job := range associations {
		fmt.Printf("%s\t%s\n", name, job)
	}
}

func printAges(ages map[string]int) {
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
