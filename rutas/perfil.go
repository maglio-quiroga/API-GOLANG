package rutas

import (
	"html/template"
	"net/http"
	"strconv"

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

func seq(start, end int) []int {
	var s []int
	for i := start; i <= end; i++ {
		s = append(s, i)
	}
	return s
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func Productos(w http.ResponseWriter, r *http.Request) {
	PaginaStr := r.URL.Query().Get("page")
	if PaginaStr == "" {
		PaginaStr = "1"
	}
	pagina, err := strconv.Atoi(PaginaStr)
	if err != nil || pagina < 1 {
		pagina = 1
	}
	const RegistroCant = 3
	alternador := (pagina - 1) * RegistroCant
	var productos []modelos.Producto

	resp := db.Database.Limit(RegistroCant).Offset(alternador).Find(&productos)
	if resp.Error != nil {
		panic(resp.Error)
	}
	var ProductosMax int64
	db.Database.Model(&modelos.Producto{}).Count(&ProductosMax)
	TotalPaginas := (int(ProductosMax) + RegistroCant - 1) / RegistroCant
	funcMap := template.FuncMap{
		"seq": seq,
		"add": add,
		"sub": sub,
	}
	data := map[string]interface{}{
		"Productos":    productos,
		"PaginasTotal": TotalPaginas,
		"PaginaActual": pagina,
	}

	t, err := template.New("p-productos.html").Funcs(funcMap).ParseFiles("templates/p-productos.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, data)
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
