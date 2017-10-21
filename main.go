package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jeffprestes/go-webserver-example-bluemix/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})
	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	fmt.Println("Estou escutando, comandante...")
	port := os.Getenv("VCAP_APP_PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
