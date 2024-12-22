package main

import (
	"log"
	"net/http"

	"POO2/db"
	"POO2/handlers"
)

func main() {
	db.InitDB("localhost", "postgres", "******", "postgres", 5432)
	defer db.DB.Close()

	http.HandleFunc("/", handlers.InicioHandler)
	http.HandleFunc("/autor", handlers.AutorFormHandler)
	http.HandleFunc("/submitAutor", handlers.EnvioAutorHandler)
	http.HandleFunc("/libro", handlers.LibroFormHandler)
	http.HandleFunc("/submitLibro", handlers.EnvioLibroHandler)
	http.HandleFunc("/eliminacion", handlers.MostrarFormularioEliminar)
	http.HandleFunc("/procesarEliminacion", handlers.ProcesarEliminacion)

	log.Println("Servidor corriendo en http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
