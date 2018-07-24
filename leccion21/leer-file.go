package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Contenido del fichero")

	texto, err := ioutil.ReadFile("file.txt")
	verError(err)

	fmt.Println(string(texto))
}

func verError(e error) {
	if e != nil {
		panic(e)
	}
}
