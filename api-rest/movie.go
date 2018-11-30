package main

type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type Movies []Movie
