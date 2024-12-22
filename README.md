Proyecto Final V 2.0.0
Utilizacion de 
  - Postgres
  - Dbeaver
  - Golang
  - Html

Para el correccto funcionamiento de este programa se requiere una base de datos de postgres.
EN LA SIGUIENT ELINEA INSERTO EL SQL PARA LA GENERACION DE LA BASE DE DATOS CON SUS RELACIONES.

CREATE TABLE autor (
  id SERIAL PRIMARY KEY,
  nombre VARCHAR(255) NOT NULL
);

CREATE TABLE genero (
  id SERIAL PRIMARY KEY,
  tipo VARCHAR(255) NOT NULL
);

//Las siguientes lineas del SQL deben ejecutarse despues de haber ejecutado las lineas anteriores.

CREATE TABLE libro (
  id SERIAL PRIMARY KEY,
  titulo VARCHAR(255),
  id_autor INT REFERENCES autor(id),
  id_genero INT REFERENCES genero(id),
  fecha_publicacion DATE NOT NULL
);

RECUERDESE CAMBIAR LAS CREDENCIALES DE CONEXION A LA BDD EN CADA UNO DE LOS PUNTOS EN LOS QUE SE REQUIERA.
