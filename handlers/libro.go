package handlers

import (
	"POO2/db"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func LibroFormHandler(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "templates/libro.html", nil)
}

func EnvioLibroHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtener los valores del formulario
		titulo := strings.TrimSpace(r.FormValue("tituloLibro"))
		idAutorStr := strings.TrimSpace(r.FormValue("idAutor"))
		idGeneroStr := strings.TrimSpace(r.FormValue("idGenero"))

		// Validar que no estén vacíos
		if titulo == "" || idAutorStr == "" || idGeneroStr == "" {
			http.Error(rw, "Todos los campos son requeridos", http.StatusBadRequest)
			return
		}

		// Convertir idAutor y idGenero a enteros
		idAutor, err := strconv.Atoi(idAutorStr)
		if err != nil {
			http.Error(rw, "ID de Autor inválido", http.StatusBadRequest)
			return
		}
		idGenero, err := strconv.Atoi(idGeneroStr)
		if err != nil {
			http.Error(rw, "ID de Género inválido", http.StatusBadRequest)
			return
		}

		// Insertar en la base de datos
		_, err = db.DB.Exec(`INSERT INTO libro (titulo, id_autor, id_genero) VALUES($1, $2, $3)`, titulo, idAutor, idGenero)
		if err != nil {
			log.Printf("Error al guardar el libro: %v", err)
			http.Error(rw, "Error al guardar en la base de datos", http.StatusInternalServerError)
			return
		}

		// Mostrar confirmación
		data := map[string]interface{}{
			"titulo":   titulo,
			"idAutor":  idAutor,
			"idGenero": idGenero,
		}
		renderTemplate(rw, "templates/index.html", data)
	} else {
		http.Error(rw, "Método no permitido", http.StatusMethodNotAllowed)
	}

}
