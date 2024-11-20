/*
	@autor: Marcelo Barrera
			Alex Hidalgo
	@fecha: 7/11/2024
	@descripcion: Iniciod el proyecto de diseño de un sitema
	de gestion
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	//area de importacion de clases
	gestionBiblioteca "gestionBiblioteca/gestion"
	//Importamos el paquene que nos ayudara a conectar con la BDD
	"database/sql"
	//Importamos el driver de  Postgres
	_ "github.com/lib/pq"
)

// Main
func main() {
	db := conexionBdd()
	defer db.Close()
	menuInicio(db)

}

// Funcion que me permite navegar en un menu
func menuInicio(db *sql.DB) {
	var opcion int

	for opcion != 4 {
		//Opciones del Menu
		fmt.Println("MENU")
		fmt.Println("1. Ingresar un libro")
		fmt.Println("4. Salir")
		fmt.Println("Que accion desea hacer: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:

			reader := bufio.NewReader(os.Stdin)

			//Iniciadores
			nlibro := gestionBiblioteca.Libro{}
			nautor := gestionBiblioteca.Autor{}

			//Solicito el titulo del libro
			fmt.Println("Título del libro: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)
			nlibro.SetTitulo(titulo)
			//Solicito la fecha de publicacion del libro
			fmt.Println("Fecha de publicacion: ")
			fmt.Println("(Formato AAAA-MM-DD)")
			fecha, _ := reader.ReadString('\n')
			fecha = strings.TrimSpace(fecha)
			nlibro.SetFecha(fecha)
			//Solicito el enlace del libro
			fmt.Println("Link del Libro: ")
			archivo, _ := reader.ReadString('\n')
			archivo = strings.TrimSpace(archivo)
			nlibro.SetArchivo(archivo)

			//Solicito la ID del autor en caso de que exista
			var op int
			/*
				fmt.println (AutorListado(db))
				fmt.Println("Presione 1 y Digite el ID del autor que desea ingresar")
				fmt.println("En caso de que no conste en la lista digite 2")
				fmt.Scan(&listautores)
			*/
			switch op {
			case 1:
				fmt.Println("Ingrese la ID correspondiente al autor")

			}

			//Solicito el Nombre de Autor
			fmt.Println("Nombre del Autor: (solo nombres) ")
			nombreA, _ := reader.ReadString('\n')
			nombreA = strings.TrimSpace(nombreA)
			nautor.SetNombre(nombreA)
			//Solicito el apellido del Autor
			fmt.Println("Apellido del autor: ")
			apellidoA, _ := reader.ReadString('\n')
			apellidoA = strings.TrimSpace(apellidoA)
			nautor.SetApellido(apellidoA)

			//Imprimo la informacion sobre el libro
			/* fmt.Println(
			nlibro.GetTitulo(),
			nlibro.GetFecha(),
			nautor.GetNombre(),
			nautor.GetApellido(),
			nautor.GetIdAutor()) */

			nlibro.IngresoLibro(db)
			nautor.IngresoAutor(db)
			nautor.IDAutor(db)
			nautor.AutorListado(db)

			// prueba de impresion del ID del autor

		case 4:
			//Opcion para terminar el programa
			os.Exit(0)
		default:
			fmt.Println("Opcion no valida")
		}

	}
}

// Funcion que nos da la conexion con la BDD
func conexionBdd() *sql.DB {
	//conexion con la BDD
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "mibe2001"
		dbname   = "Biblioteca"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//Establecemos la conexion
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error al conectar con la base de datos", err)
		os.Exit(1) // el valor de 1 indica q el codigo a generado un error por lo que se usa osExit para una  salida inmediata
	}

	//Comprobacion de la conexion
	err = db.Ping()
	if err != nil {
		fmt.Println("Error al conectar con la base de datos", err)
		os.Exit(1) // el valor de 1 indica q el codigo a generado un error por lo que se usa os.Exit para una  salida inmediata
	}

	fmt.Println("Conexion a la BASE DE DATOS realizada con EXITO")
	return db
}

// Funcion para ingresar un libro
// En este caso no se solicita un ID puesto que la BDD
// se encargara de crear uno mediante parametros de identidad.
