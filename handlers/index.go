package handlers

import (
	"net/http"
)

func InicioHandler(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "templates/index.html", nil)
}
