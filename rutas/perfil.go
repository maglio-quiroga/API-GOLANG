package rutas

import (
	"html/template"
	"net/http"

	"github.com/maglio-quiroga/API-GOLANG/db"
	"github.com/maglio-quiroga/API-GOLANG/modelos"
)

func PaginaPerfil(w http.ResponseWriter, r *http.Request) {
	sesion := ObtenerSesion(r)
	var usuario modelos.Usuario
	t, err := template.ParseFiles("templates/perfil.html")
	if err != nil {
		panic(err)
	}
	registro := db.Database.Where("id = ?", sesion.Values["user_id"]).First(&usuario)
	if registro.Error != nil {
		panic(registro.Error)
	}
	datos := map[string]interface{}{
		"Nombre":   usuario.Nombre,
		"Apellido": usuario.Apellido,
		"Email":    usuario.Email,
	}
	t.Execute(w, datos)
}

func Productos(w http.ResponseWriter, r *http.Request) {
	//sesion := ObtenerSesion(r)
	//var usuario modelos.Usuario
	t, err := template.ParseFiles("templates/p-productos.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func CerrarSesion(w http.ResponseWriter, r *http.Request) {
	sesion := ObtenerSesion(r)
	sesion.Values = make(map[interface{}]interface{})
	err := sesion.Save(r, w)
	if err != nil {
		http.Error(w, "Error al guardar la sesión", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
