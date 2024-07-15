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
