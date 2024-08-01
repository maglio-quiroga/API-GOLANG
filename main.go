package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/maglio-quiroga/API-GOLANG/db"
	"github.com/maglio-quiroga/API-GOLANG/modelos"
	"github.com/maglio-quiroga/API-GOLANG/rutas"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	var eventos []modelos.Evento
	var banner []modelos.Banner
	var productos []modelos.Producto
	db.Database.Find(&banner)
	db.Database.Find(&eventos)
	db.Database.Order("id desc").Limit(3).Find(&productos)
	data := map[string]interface{}{
		"Eventos":   eventos,
		"Banner":    banner,
		"Productos": productos,
	}

	t.Execute(w, data)
}

func main() {
	db.Iniciar_conexion()
	db.Database.AutoMigrate(modelos.Usuario{})
	db.Database.AutoMigrate(modelos.Producto{})
	db.Database.AutoMigrate(modelos.Banner{})
	db.Database.AutoMigrate(modelos.Evento{})

	enrutador := mux.NewRouter()
	enrutador.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	enrutador.PathPrefix("/archivos/").Handler(http.StripPrefix("/archivos/", http.FileServer(http.Dir("./archivos/"))))
	enrutador.HandleFunc("/", Index)
	enrutador.HandleFunc("/Iniciar", rutas.Iniciar_sesion).Methods("GET", "POST")
	enrutador.HandleFunc("/Registrarse", rutas.Registrarse).Methods("GET", "POST")
	enrutador.Handle("/salir", rutas.AutenticarSoggaShop(http.HandlerFunc(rutas.CerrarSesion)))
	enrutador.Handle("/perfil", rutas.AutenticarSoggaShop(http.HandlerFunc(rutas.PaginaPerfil)))
	enrutador.Handle("/productos", rutas.AutenticarSoggaShop(http.HandlerFunc(rutas.Productos)))
	enrutador.Handle("/adm-inicio", rutas.AutenticarAdm(http.HandlerFunc(rutas.InicioAdm)))
	enrutador.Handle("/form-banner", rutas.AutenticarAdm(http.HandlerFunc(rutas.SubirArchivo)))
	enrutador.Handle("/form-eventos", rutas.AutenticarAdm(http.HandlerFunc(rutas.CrearEventos)))
	enrutador.Handle("/eliminar-banner", rutas.AutenticarAdm(http.HandlerFunc(rutas.EliminarBanner)))
	enrutador.Handle("/eliminar-evento", rutas.AutenticarAdm(http.HandlerFunc(rutas.EliminarEvento)))

	http.ListenAndServe(":3000", enrutador)

}
