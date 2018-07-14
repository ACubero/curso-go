package main

import (
	"fmt"
)

func main() {
	//Definición de array forma 1
	var vtestamento [3]string
	vtestamento[0] = "Génesis"
	vtestamento[1] = "Éxodo"
	vtestamento[2] = "Levítico"

	fmt.Println(vtestamento)

	//Definición de arrays forma 2
	ntestamento := [3]string{"Mateo", "Marcos", "Lucas"}

	fmt.Println(ntestamento)
	fmt.Println("Final")
}
