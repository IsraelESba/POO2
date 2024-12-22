package db

import (
	"POO2/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(host, user, password, dbname string, port int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa.")
}

// ObtenerLibros obtiene todos los libros desde la base de datos
func ObtenerLibros() ([]models.Libro, error) {
	// Realizamos la consulta a la base de datos
	rows, err := DB.Query("SELECT id, titulo FROM libro")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []models.Libro
	for rows.Next() {
		var libro models.Libro
		if err := rows.Scan(&libro.ID, &libro.Titulo); err != nil { //cambias por $autor.ID ...
			return nil, err
		}
		libros = append(libros, libro)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return libros, nil
}

//Obtener Autores desde la base de datos
