package handlers

import (
	"POO2/db"
	"Prueba/models"
	"html/template"
	"net/http"
)

// EliminarHandler maneja la solicitud para la página de eliminación de libros
func EliminarHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Conecta a la base de datos
		dbConn, err := db.ConnectDB()
		if err != nil {
			http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
			return
		}
		defer dbConn.Close()

		// Obtiene la lista de libros desde la base de datos
		rows, err := dbConn.Query("SELECT id, titulo, autor FROM libros")
		if err != nil {
			http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var libros []models.Libro
		for rows.Next() {
			var libro models.Libro
			if err := rows.Scan(&libro.ID, &libro.Titulo); err != nil {
				http.Error(w, "Error al procesar los datos de la base de datos", http.StatusInternalServerError)
				return
			}
			libros = append(libros, libro)
		}

		// Renderiza la plantilla HTML con la lista de libros
		if err := tmpl.ExecuteTemplate(w, "eliminacion.html", libros); err != nil {
			http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		}
	}
}
