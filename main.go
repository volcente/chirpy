package main

import (
	"log"
	"net/http"
)

func main() {
	filepathRoot := "."
	port := "8080"

	serverMux := http.NewServeMux()
	serverMux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	server := http.Server{
		Addr:    port,
		Handler: serverMux,
	}

	log.Printf("server started at http://localhost:%s, serving files from :%s\n", port, filepathRoot)
	log.Fatal(server.ListenAndServe())
}
