package main

import (
	"fmt"
)

func main() {
	// Declaramos dos variables como punteros enteros
	var x, y *int

	// Declaramos una variable con un valor de 5 entero
	entero := 5

	// Con el & recogemos la dirección de la memoria de la variable entero y se lo pasamos a x,y
	x = &entero
	y = &entero

	//Pintamos, nos muestra la dirección de memoria de x,y
	fmt.Println(x)
	fmt.Println(y)

	// Cambiamos el valor que tiene x, antes era 5, porque le hemos asignado entero
	// Ahora con el *, lo que hacemos es cambiar el VALOR que había asignado en la direccion de memoria que tenía X
	*x = 53

	//Como la dirección de x,y es la misma si pintamos el valor de ambas con el *, nos aparece el mismo valor
	//Si pintamos la variable entero, vale lo mismo que x,y porque hemos modificado la dirección de la memoria
	//donde se almacenaba x,y,entero
	fmt.Println(*x)
	fmt.Println(*y)
	fmt.Println(entero)
}
