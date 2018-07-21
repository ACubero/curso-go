package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("** Conversor de medidas ***")
	// Recogemos el parámetro1 y parametro 2
	fmt.Println("Introduce convertir " + os.Args[1] + " a " + os.Args[2])

}
