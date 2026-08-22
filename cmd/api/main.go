package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerInit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Inicio de aplicacion")
}

func main() {
	const port int = 8080
	var addr string = fmt.Sprintf(":%d", port)

	http.HandleFunc("/", handlerInit)

	log.Println("Servidor corriendo en puerto : ", port)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalln("Ocurrio un error al iniciar servidor", err)
	}

}
