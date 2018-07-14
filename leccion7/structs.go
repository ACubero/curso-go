package main

import (
	"fmt"
)

type Estudiante struct {
	nombre    string
	apellidos string
	edad      int
	hombre    bool
}

func main() {
	var alumno1 = Estudiante{
		nombre:    "Alex",
		apellidos: "Cubero Arrizabalaga ",
		edad:      22,
		hombre:    true}
	var alumno2 = Estudiante{"Juan", "Primo Rivera", 23, true}
	fmt.Println(alumno1)
	fmt.Println(alumno2)
	fmt.Println(alumno2.nombre)
	fmt.Println(alumno2.apellidos)
	fmt.Println(alumno2.edad)
	fmt.Println(alumno2.hombre)
}
