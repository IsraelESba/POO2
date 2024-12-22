package handlers

import (
	"POO2/models"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
)

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
	rows, err := db.Query("SELECT id, titulo FROM libro ORDER BY id")
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

// Eliminar un libro por ID
func eliminarLibro(id int) error {
	connStr := "user=postgres dbname=postgres password=mibe2001 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM libro WHERE id = $1", id)
	return err
}

// Manejador para mostrar el formulario
func MostrarFormularioEliminar(w http.ResponseWriter, r *http.Request) {
	libros, err := obtenerLibros()
	if err != nil {
		http.Error(w, "No se pudo obtener la lista de libros", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/eliminacion.html")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, libros)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
	}
}

// Manejador para procesar la eliminación
func ProcesarEliminacion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := eliminarLibro(id); err != nil {
		http.Error(w, "Error al eliminar el libro", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
