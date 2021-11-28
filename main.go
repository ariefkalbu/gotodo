package main

import (
	"goweb/controller"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controller.Login)
	mux.HandleFunc("/login", controller.ProcessLogin)
	mux.HandleFunc("/signup", controller.ProcessSignup)

	css := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", css))

	img := http.FileServer(http.Dir("assets/img"))
	mux.Handle("/img/", http.StripPrefix("/img", img))

	log.Print("localhost:71")

	err := http.ListenAndServe(":71", mux)

	log.Println(err)
}
