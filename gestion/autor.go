package gestionBiblioteca

import (
	_ "bufio"
	"database/sql"
	"fmt"
	_ "os"
	_ "strings"
)

//creacion del objeto

type Autor struct {
	nombre string
	IdA    int
}

// zona de set's
func (a *Autor) SetNombre(nombre string) {
	a.nombre = nombre
}

//zona de get's

func (a *Autor) GetNombre() string {
	return a.nombre
}

func (a *Autor) GetIdAutor() int {
	return a.IdA
}

//hacemos el ingreso del autor a la bdd

func (a *Autor) IngresoAutor(db *sql.DB) {

	// Insertar un nuevo autor
	_, err := db.Exec("INSERT INTO autor (nombre) VALUES ($1)",
		a.nombre)
	if err != nil {
		fmt.Println(err)
	}
}

// obtener la ID del autor nuuevo ingresado

func (a *Autor) IDAutor(db *sql.DB) {
	//recuperar el ID del autor

	_, err := db.Exec("SELECT id FROM autor", a.IdA)
	if err != nil {
		fmt.Println(err)
	}
}

func (a *Autor) AutorListado(db *sql.DB) {
	//mostramos la lista de autores

	_, err := db.Exec("SELECT * FROM autor", a.IdA)
	if err != nil {
		fmt.Println(err)
	}

}

//mostrar los libros pertenecientes a cada autor
