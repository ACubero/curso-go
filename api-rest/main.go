package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true) // URLS amigables
	router.HandleFunc("/", Index)
	router.HandleFunc("/contacto", Contact)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Estamos en la página de contactos")
}
