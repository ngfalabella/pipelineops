package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerInit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Inicio de aplicacion")
}

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "OK")
	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}

func main() {
	const port int = 8080
	var addr string = fmt.Sprintf(":%d", port)

	http.HandleFunc("/", handlerInit)
	http.HandleFunc("/healthz", handlerHealthz)

	log.Println("Servidor corriendo en puerto : ", port)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalln("Ocurrio un error al iniciar servidor", err)
	}

}
