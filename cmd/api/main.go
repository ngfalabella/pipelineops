package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerInit)
	mux.HandleFunc("/healthz", handlerHealthz)

	log.Println("Servidor corriendo en puerto : ", port)

	servidor := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,  //tiempo maximo para leer la request
		WriteTimeout: 10 * time.Second, //tiempo maximo para escribir la respuesta
		IdleTimeout:  60 * time.Second, //tiempo que vamos a permitir que la conexion este abierta sin hacer nada
	}

	err := servidor.ListenAndServe()

	if err != nil {
		log.Fatalln("Ocurrio un error al iniciar servidor", err)
	}

}
