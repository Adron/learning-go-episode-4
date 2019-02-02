package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, let's talk composite types.")

	basketOfStuff := [3]string{"The first string","second","This string."}
	var zeeValues [2]int

	for i, v := range basketOfStuff {
		fmt.Printf("Value %d: %s\n", i, v)
	}

	fmt.Println(zeeValues)

	if zeeValues[0] == zeeValues[1] {
		fmt.Println("The values are the same, this doesn't instantiate like the `new` keyword.")
	} else {
		fmt.Println("The way go appears to instantiate unset variable values, such as in this array is like the `new` keyword instantiation.")
	}

	zeeValues[0] = 1 + 52 * 3
	zeeValues[1] = 9

	fmt.Println(zeeValues[len(zeeValues) - 1])

	type Currency int

	const (
		USD Currency = iota
		CAN
		EUR
		GBP
		JPY
		NOK
		SEK
		DKK
	)

	symbol := [...]string{USD: "$", CAN: "$",  EUR: "€", GBP: "£", JPY:"¥", NOK:"kr", SEK:"kr",DKK:"kr"}

	fmt.Println(EUR, symbol[EUR])

	fmt.Println(JPY, symbol[JPY])

	r := [...]int{99: -1}

	r[36] = 425
	r[42] = 42

	fmt.Println(r[36] + r[42])
	fmt.Println(strconv.Itoa(r[36]))

	months := [...]string{1: "January", 2:"February", 3: "March", 4:"April", 12:"December"}

	for _, s := range months {
		fmt.Printf("The month: %s\n", s)
	}

	var runes []rune
	for _, r := range "Language: 走" {
		runes = append(runes, r)
	}
	fmt.Printf("%q \n", runes)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

}

func appendInt(x []int, i int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2* len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	return z
}
