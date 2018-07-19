package main

import "fmt"

func main() {
	anos := 18
	// !=, ==, &&, || >= <=
	if (anos > 18 && anos == 18) || anos == 17 {
		fmt.Println("Eres mayor de edad")
	} else if anos == 16 {
		fmt.Println("Eres MENOR de edad")
	}
}
