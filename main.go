package main

import (
	"log"
	"net/http"

	"Prueba/db"
	"Prueba/handlers"
)

func main() {
	db.InitDB("localhost", "postgres", "12345", "postgres", 5432)
	defer db.DB.Close()

	http.HandleFunc("/", handlers.AutorFormHandler)
	http.HandleFunc("/submitAutor", handlers.EnvioAutorHandler)
	http.HandleFunc("/libro", handlers.LibroFormHandler)
	http.HandleFunc("/submitLibro", handlers.EnvioLibroHandler)

	log.Println("Servidor corriendo en localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
