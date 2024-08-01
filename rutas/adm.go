package rutas

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/maglio-quiroga/API-GOLANG/db"
	"github.com/maglio-quiroga/API-GOLANG/modelos"
)

func InicioAdm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/adm-inicio.html")
	if err != nil {
		panic(err)
	}
	var banner []modelos.Banner
	var eventos []modelos.Evento
	db.Database.Find(&banner)
	db.Database.Find(&eventos)
	data := map[string]interface{}{
		"Banner":  banner,
		"Eventos": eventos,
	}
	t.Execute(w, data)
}

func SubirArchivo(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/adm-inicio.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		r.ParseMultipartForm(2000)
		file, fileInfo, _ := r.FormFile("imgfile")
		titulo := r.Form["tituloimg"][0]
		f, err := os.OpenFile("./archivos/"+fileInfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		directorio := "./archivos/" + fileInfo.Filename
		CamposVacios := false
		ProblemaArchivo := false
		if err != nil {
			ProblemaArchivo = true
			panic(err)
		}
		if titulo == "" || fileInfo.Filename == "" {
			CamposVacios = true
		}
		data := map[string]interface{}{
			"CamposVacios":    CamposVacios,
			"ProblemaArchivo": ProblemaArchivo,
		}
		if titulo != "" && fileInfo.Filename != "" {
			registro := &modelos.Banner{Titulo: titulo, Imagen: directorio}
			resp := db.Database.Create(registro)
			if resp.Error != nil {
				panic(resp.Error)
			} else {
				fmt.Println("registro insertado")
				http.Redirect(w, r, "/adm-inicio", http.StatusFound)
			}
		}
		t, err := template.ParseFiles("templates/adm-inicio.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, data)

		defer f.Close()
		io.Copy(f, file)
	}
}

func CrearEventos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/adm-inicio.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		nombre := r.Form["nombre"][0]
		desc := r.Form["desc"][0]
		fecha := r.Form["fecha"][0]
		CamposVacios := false
		if nombre == "" || desc == "" || fecha == "" {
			CamposVacios = true
		}
		data := map[string]interface{}{
			"CamposVacios": CamposVacios,
		}
		if nombre != "" && desc != "" && fecha != "" {
			registro := &modelos.Evento{Nombre: nombre, Descripcion: desc, FechaTermino: fecha}
			resp := db.Database.Create(registro)
			if resp.Error != nil {
				panic(resp.Error)
			} else {
				fmt.Println("Evento Insertado")
				http.Redirect(w, r, "/adm-inicio", http.StatusFound)
			}
		}
		t, err := template.ParseFiles("templates/adm-inicio.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, data)
	}
}

func EliminarBanner(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var banner modelos.Banner
		id := r.Form["idbanner"][0]
		db.Database.Where("id = ?", id).Delete(&banner)
		http.Redirect(w, r, "/adm-inicio", http.StatusFound)
		//faltaria agregar la eliminacion del archivo y no solo la del registro
	}
}

func EliminarEvento(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var evento modelos.Evento
		id := r.Form["idevento"][0]
		db.Database.Where("id = ?", id).Delete(&evento)
		http.Redirect(w, r, "/adm-inicio", http.StatusFound)
	}
}

func CerrarAdm(w http.ResponseWriter, r *http.Request) {
	sesion := ObtenerSesion(r)
	sesion.Values = make(map[interface{}]interface{})
	err := sesion.Save(r, w)
	if err != nil {
		http.Error(w, "Error al guardar la sesión", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
