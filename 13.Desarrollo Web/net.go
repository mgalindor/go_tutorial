package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la página de inicio")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la página Acerca de")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Servidor escuchando en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}
