package handlers

import (
	"POO2/db"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func LibroFormHandler(rw http.ResponseWriter, r *http.Request) {
	// Obtener los autores desde la base de datos
	rows, err := db.DB.Query("SELECT id, nombre FROM autor")
	if err != nil {
		log.Printf("Error al obtener autores: %v", err)
		http.Error(rw, "Error al obtener autores desde la base de datos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var autores []map[string]interface{}
	for rows.Next() {
		var id int
		var nombre string
		if err := rows.Scan(&id, &nombre); err != nil {
			log.Printf("Error al leer autor: %v", err)
			continue
		}
		autores = append(autores, map[string]interface{}{
			"id":     id,
			"nombre": nombre,
		})
	}

	// Obtener los géneros desde la base de datos
	rowsGeneros, err := db.DB.Query("SELECT id, tipo FROM genero")
	if err != nil {
		log.Printf("Error al obtener géneros: %v", err)
		http.Error(rw, "Error al obtener géneros desde la base de datos", http.StatusInternalServerError)
		return
	}
	defer rowsGeneros.Close()

	var generos []map[string]interface{}
	for rowsGeneros.Next() {
		var id int
		var tipo string
		if err := rowsGeneros.Scan(&id, &tipo); err != nil {
			log.Printf("Error al leer género: %v", err)
			continue
		}
		generos = append(generos, map[string]interface{}{
			"id":   id,
			"tipo": tipo,
		})
	}

	// Pasar los autores y géneros a la plantilla
	data := map[string]interface{}{
		"autores": autores,
		"generos": generos,
	}

	renderTemplate(rw, "templates/libro.html", data)
}

func EnvioLibroHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Obtener los valores del formulario
		titulo := strings.TrimSpace(r.FormValue("tituloLibro"))
		idAutorStr := strings.TrimSpace(r.FormValue("idAutor"))
		idGeneroStr := strings.TrimSpace(r.FormValue("idGenero"))
		fechaPublicacion := strings.TrimSpace(r.FormValue("fechaPublicacion"))

		// Validar que no estén vacíos
		if titulo == "" || idAutorStr == "" || idGeneroStr == "" || fechaPublicacion == "" {
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
		_, err = db.DB.Exec(`INSERT INTO libro (titulo, id_autor, id_genero, fecha_publicacion) VALUES($1, $2, $3, $4)`, titulo, idAutor, idGenero, fechaPublicacion)
		if err != nil {
			log.Printf("Error al guardar el libro: %v", err)
			http.Error(rw, "Error al guardar en la base de datos", http.StatusInternalServerError)
			return
		}

		// Mostrar confirmación
		data := map[string]interface{}{
			"titulo":           titulo,
			"idAutor":          idAutor,
			"idGenero":         idGenero,
			"fechaPublicacion": fechaPublicacion,
		}
		renderTemplate(rw, "templates/index.html", data)
	} else {
		http.Error(rw, "Método no permitido", http.StatusMethodNotAllowed)
	}

}
