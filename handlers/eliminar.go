package handlers

import (
	"POO2/models"
	"database/sql"
	"html/template"
	"net/http"
)

type Book struct {
	ID     int
	Titulo string
}

// Conexión a la base de datos PostgreSQL
func obtenerLibros() ([]models.Libro, error) {
	// Aquí va la conexión a tu base de datos PostgreSQL
	connStr := "user=postgres dbname=postgres password=mibe2001 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Consultamos los libros
	rows, err := db.Query("SELECT id, titulo FROM libro")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []models.Libro
	for rows.Next() {
		var libro models.Libro
		if err := rows.Scan(&libro.ID, &libro.Titulo); err != nil {
			return nil, err
		}
		libros = append(libros, libro)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return libros, nil
}

// Manejador HTTP que renderiza la página HTML
func HandlerEliminar(w http.ResponseWriter, r *http.Request) {
	libros, err := obtenerLibros()
	if err != nil {
		http.Error(w, "No se pudo obtener los libros", http.StatusInternalServerError)
		return
	}

	// Cargar la plantilla HTML desde el archivo
	tmpl, err := template.ParseFiles("templates/eliminacion.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	// Ejecutar la plantilla HTML con los datos de los libros
	err = tmpl.Execute(w, libros)
	if err != nil {
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
}
