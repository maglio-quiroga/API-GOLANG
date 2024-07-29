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
	t.Execute(w, nil)
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
			}
		}
		t, err := template.ParseFiles("templates/adm-inicio.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, data)
	}
}
