package rutas

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/gorilla/sessions"
	"github.com/maglio-quiroga/API-GOLANG/db"
	"github.com/maglio-quiroga/API-GOLANG/modelos"
)

var almacenamiento = sessions.NewCookieStore([]byte("clave-super-secreta"))

func ObtenerSesion(r *http.Request) *sessions.Session {
	session, _ := almacenamiento.Get(r, "session-name")
	return session
}

func AutenticarSoggaShop(sig http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sesion := ObtenerSesion(r)
		if sesion.Values["user_id"] == nil {
			http.Redirect(w, r, "/Iniciar", http.StatusFound)
			return
		}
		sig.ServeHTTP(w, r)
	})
}

func Iniciar_sesion(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/iniciar.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		email := r.Form["correo"][0]
		clave := r.Form["clave"][0]
		hash := md5.New()
		hash.Write([]byte(clave))
		ClaveEncriptada := hex.EncodeToString(hash.Sum(nil))
		var usuario modelos.Usuario
		CamposVacios := false
		NoExiste := false
		if email == "" || clave == "" {
			CamposVacios = true
		}
		if err := db.Database.Where("email = ?", email).First(&usuario).Error; err != nil {
			NoExiste = true
		}
		data := map[string]interface{}{
			"CamposVacios": CamposVacios,
			"NoExiste":     NoExiste,
		}
		if data["CamposVacios"] == false && data["NoExiste"] == false {
			if err := db.Database.Where("email = ? AND clave = ?", email, ClaveEncriptada).First(&usuario).Error; err == nil {

				sesion := ObtenerSesion(r)
				sesion.Values["user_id"] = usuario.ID
				sesion.Save(r, w)

				http.Redirect(w, r, "/perfil", http.StatusFound)
				return
			}
		}
		t, err := template.ParseFiles("templates/iniciar.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, data)
	}
}

func Registrarse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/registrarse.html")
		if err != nil {
			fmt.Println("El nombre del archivo o el directorio son incorrectos")
			panic(err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		nombre := r.Form["nombre"][0]
		apellido := r.Form["apellido"][0]
		email := r.Form["correo"][0]
		clave1 := r.Form["clave1"][0]
		clave2 := r.Form["clave2"][0]
		CamposVacios := false
		NombreIncorrectos := false
		ExisteUsuario := false
		letrasRegexp := regexp.MustCompile(`^[a-zA-Z]+$`)
		var usuario modelos.Usuario

		if nombre == "" {
			CamposVacios = true
		} else if !letrasRegexp.MatchString(nombre) {
			NombreIncorrectos = true
		}
		if apellido == "" {
			CamposVacios = true
		} else if !letrasRegexp.MatchString(apellido) {
			NombreIncorrectos = true
		}
		if email == "" {
			CamposVacios = true
		} else if err := db.Database.Where("email = ?", email).First(&usuario).Error; err == nil {
			ExisteUsuario = true
		}
		if clave1 == "" || clave2 == "" {
			CamposVacios = true
		}

		data := map[string]interface{}{
			"Nombre":             nombre,
			"Apellido":           apellido,
			"Email":              email,
			"Clave1":             clave1,
			"Clave2":             clave2,
			"ClaveIncorrecta":    clave1 != clave2,
			"CamposVacios":       CamposVacios,
			"NombresIncorrectos": NombreIncorrectos,
			"ExisteUsuario":      ExisteUsuario,
		}
		hash := md5.New()
		hash.Write([]byte(clave1))
		ClaveEncriptada := hex.EncodeToString(hash.Sum(nil))
		if data["ClaveIncorrecta"] == false && data["CamposVacios"] == false && data["NombresIncorrectos"] == false && data["ExisteUsuario"] == false {
			fmt.Println("Las credenciales son correctas")
			registro := &modelos.Usuario{Nombre: nombre, Apellido: apellido, Email: email, Clave: ClaveEncriptada}
			resp := db.Database.Create(registro)
			if resp.Error != nil {
				fmt.Println("ERROR")
				panic(resp.Error)
			} else {
				fmt.Println("Usuario insertado")
			}
		}
		t, err := template.ParseFiles("templates/registrarse.html")
		if err != nil {
			fmt.Println("El nombre del archivo o el directorio son incorrectos")
			panic(err)
		}
		t.Execute(w, data)
	}
}
