package main

import "fmt"

func main() {
	fmt.Println(saberEdad())
}

func saberEdad() (texto string, edad int) {
	texto = "Esto sería un texto"
	edad = 10
	return texto, edad
}
