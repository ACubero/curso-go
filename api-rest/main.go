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
	router.HandleFunc("/movies", MovieList)
	router.HandleFunc("/movie/{id}", MovieShow)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}
func MovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listado de películas")
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprintf(w, "La película que quiere ver %s", movieId)
}
