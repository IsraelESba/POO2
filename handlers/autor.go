package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"Prueba/db"
)

func AutorFormHandler(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "templates/autor.html", nil)
}

func EnvioAutorHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := strings.TrimSpace(r.FormValue("nombreAutor"))
		if nombre == "" {
			http.Error(rw, "El nombre es requerido", http.StatusBadRequest)
			return
		}
		_, err := db.DB.Exec(`INSERT INTO autor (nombre) VALUES($1)`, nombre)
		if err != nil {
			log.Printf("Error al guardar el autor: %v", err)
			http.Error(rw, "Error al guardar en la base de datos", http.StatusInternalServerError)
			return
		}
		http.Redirect(rw, r, "/libro", http.StatusSeeOther)
	} else {
		http.Error(rw, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func renderTemplate(rw http.ResponseWriter, filepath string, data interface{}) {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(rw, "Error al cargar la plantilla", http.StatusInternalServerError)
		log.Printf("Error al cargar plantilla: %v", err)
		return
	}
	tmpl.Execute(rw, data)
}
