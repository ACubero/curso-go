package main

import (
	"encoding/json"
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
	movies := Movies{
		Movie{"La Guerra de las Galaxias", 1976, "Lucas Films"},
		Movie{"El Imperio contraataca", 1978, "Lucas Films"},
		Movie{"El retorno del jedi", 1992, "Lucas SkyWker"},
	}
	//fmt.Fprintf(w, "Listado de películas")
	json.NewEncoder(w).Encode(movies)
}
func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprintf(w, "La película que quiere ver %s", movieId)
}
