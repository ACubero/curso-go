package main

import "fmt"

func main() {
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Alex"

	var apellidos string = "Cubero ALCUAR"
	apellidos = "Tomás Andresa"

	pais := "España"
	pais = "Colombia"

	fmt.Println("Hola mundo desde Go  con " + nombre + " " + apellidos + " " + pais)
	fmt.Println(suma)
	fmt.Println(resta)
}
