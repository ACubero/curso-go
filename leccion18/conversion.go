package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("** Conversor de medidas ***")
	// Recogemos el parámetro1 y parametro 2

	if len(os.Args) < 2 {
		fmt.Println("No ha indicado los parámetros necesarios")
		os.Exit(0)
	}

	param1 := os.Args[1]
	param2, _ := strconv.Atoi(os.Args[2])
	conversion := 0
	if param1 == "m" {
		conversion = param2 * 10
		fmt.Println(strconv.Itoa(param2) + " metros son  " + strconv.Itoa(conversion) + " cms")
	}

}
