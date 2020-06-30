package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("SERVING_PORT")
	appMode := os.Getenv("APP_MODE")
	appVersion := os.Getenv("APP_VERSION")

	info := fmt.Sprintf("Servidor %s versão %s escutando na porta %s!", appMode, appVersion, port)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ola Mundo!")
	})

	fmt.Println(info)

	http.ListenAndServe(port, r)
}
