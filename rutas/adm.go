package rutas

import (
	"net/http"
	"text/template"
)

func InicioAdm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/adm-inicio.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
