package main

import "fmt"

func main() {

	paramIlimitados("primero", "segundo", "tercero", "cuarto")
}

//El simbolo de los tres puntos ... sirve para decir que se enviarán un número indeterminado
// de parámetros
func paramIlimitados(parametros ...string) {

	//El for necesita una variable, cuando ponemos el subrallado bajo _ sustituimos esa variable
	// Este for lo que hace es recorrer todo lo que hay guardado en parámetros y cada uno de los
	// contenidos que tiene se pasa a la variable parametros
	for _, parametro := range parametros {
		fmt.Println(parametro)
	}
}
