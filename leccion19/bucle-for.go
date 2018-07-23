package main

import "fmt"

func main() {

	//Ejemplo 1
	for i := 1; i <= 15; i++ {
		fmt.Println(i, "El número actual")
	}

	//Ejemplo 2
	max := 10
	for i := 1; i <= max; i++ {
		fmt.Println(i, "El número actual")
	}

	//Ejemplo 3
	animales := []string{"lobro", "cordero", "jabalí", "orangután", "perro", "gato", "loro"}
	for i := 0; i < len(animales); i++ {
		fmt.Println(animales[i] + " es un animal")
	}

	//Ejemplo 4
	for _, animal := range animales {
		fmt.Println(animal)

	}
}
