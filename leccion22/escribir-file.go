package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Contenido del fichero")
	//nuevo_texto := []byte(os.Args[1])
	//escribir := ioutil.WriteFile("file.txt", nuevo_texto, 0777)
	nuevo_texto := os.Args[1] + "\n"

	fichero, err := os.OpenFile("file.txt", os.O_APPEND, 0777)
	verError(err)

	escribir, err := fichero.WriteString(nuevo_texto)
	log.Println(escribir)
	verError(err)

	texto, err := ioutil.ReadFile("file.txt")
	verError(err)

	fmt.Println(string(texto))
}

func verError(e error) {
	if e != nil {
		panic(e)
	}
}
