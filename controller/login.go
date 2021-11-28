package controller

import (
	"goweb/model"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "login.html"), path.Join("views", "__layout.html"))

	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		http.NotFound(w, r)
		return
	}
}

func ProcessLogin(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, "Data tidak valid", http.StatusInternalServerError)
			return
		}

		data := model.Session{
			Username:  "ariefkalbu",
			Name:      "Arief Kalbu",
			Email:     r.Form.Get("email"),
			Logindate: time.Now(),
		}

		tmpl, err := template.ParseFiles(path.Join("views", "home.html"), path.Join("views", "__layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Home tidak ditemukan", http.StatusBadRequest)
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			log.Println(err)
			http.Error(w, "Data tidak ditemukan", http.StatusBadRequest)
			return
		}

		return
	}

	http.Error(w, "Eror saat menerima data", http.StatusBadRequest)
}

func ProcessSignup(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		err := r.ParseForm()

		if err != nil {
			log.Println(err)
			http.Error(w, "Data tidak valid", http.StatusInternalServerError)
			return
		}

		data := model.Session{
			Username:  "ariefkalbu",
			Name:      r.Form.Get("name"),
			Email:     r.Form.Get("email"),
			Logindate: time.Now(),
		}

		tmpl, err := template.ParseFiles(path.Join("views", "home.html"), path.Join("views", "__layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Home tidak ditemukan", http.StatusBadRequest)
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			log.Println(err)
			http.Error(w, "Data tidak ditemukan", http.StatusBadRequest)
			return
		}

		return
	}

	http.Error(w, "Eror saat menerima data", http.StatusBadRequest)
}
