package handlers

import (
	"log"
	"net/http"
	"strings"

	"Prueba/db"
)

func LibroFormHandler(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "templates/libro.html", nil)
}

func EnvioLibroHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		titulo := strings.TrimSpace(r.FormValue("tituloLibro"))
		if titulo == "" {
			http.Error(rw, "El título es requerido", http.StatusBadRequest)
			return
		}
		_, err := db.DB.Exec(`INSERT INTO libro (titulo) VALUES($1)`, titulo)
		if err != nil {
			log.Printf("Error al guardar el libro: %v", err)
			http.Error(rw, "Error al guardar en la base de datos", http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{"titulo": titulo}
		renderTemplate(rw, "templates/index.html", data)
	} else {
		http.Error(rw, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
