package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/maglio-quiroga/API-GOLANG/db"
	"github.com/maglio-quiroga/API-GOLANG/modelos"
	"github.com/maglio-quiroga/API-GOLANG/rutas"
)

func Index(w http.ResponseWriter, r *http.Request) {
	vista, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("El nombre del archivo o el directorio son incorrectos")
		panic(err)
	}
	vista.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	db.Iniciar_conexion()
	db.Database.AutoMigrate(modelos.Usuario{})

	enrutador := mux.NewRouter()
	enrutador.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	enrutador.HandleFunc("/", Index)
	enrutador.HandleFunc("/Iniciar", rutas.Iniciar_sesion).Methods("GET", "POST")
	enrutador.HandleFunc("/Registrarse", rutas.Registrarse).Methods("GET", "POST")
	enrutador.Handle("/perfil", rutas.AutenticarSoggaShop(http.HandlerFunc(rutas.PaginaPerfil)))

	http.ListenAndServe(":3000", enrutador)

}
