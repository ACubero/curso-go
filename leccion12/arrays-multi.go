package main

import (
	"fmt"
)

func main() {
	// Array multidimensional
	var vtestamento [3][2]string
	vtestamento[0][0] = "Génesis"
	vtestamento[0][1] = "Moisés"
	vtestamento[1][0] = "Éxodo"
	vtestamento[1][1] = "Moisés"
	vtestamento[2][0] = "Levítico"
	vtestamento[2][1] = "Moisés"

	//Se muestra todo el array
	fmt.Println(vtestamento)
	//Muestra un solo elemento
	fmt.Println(vtestamento[0][1])

}
