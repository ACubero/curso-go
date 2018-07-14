package main

import (
	"fmt"
)

func main() {
	fmt.Println(saberPrecioPedido(18, 89.2))
}
func saberPrecioPedido(numArticulos float32, preciounitario float32) (string, float32) {
	// Esto que definimos a continuación es  el closures
	precio := func() float32 {
		return numArticulos * preciounitario
	}

	return "El precio del pedido es:", precio()
}
