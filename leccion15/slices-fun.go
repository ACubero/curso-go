package main

import (
	"fmt"
)

func main() {
	//Un slice es un array pero sin definir el número de registros fijos que habrá
	ejemploSlice := []string{
		"primero",
		"segundo",
		"tercero",
		"cuarto",
		"quinto"}
	ejemploSlice = append(ejemploSlice, "sexto")
	ejemploSlice = append(ejemploSlice, "septimo")

	fmt.Println(len(ejemploSlice))
	fmt.Println(ejemploSlice[0:3])

}
