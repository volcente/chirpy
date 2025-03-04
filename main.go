package main

import (
	"log"
	"net/http"
)

func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Something went wrong while writing the response: %v", err)
	}

}

func main() {
	filepathRoot := "."
	port := ":8080"

	serverMux := http.NewServeMux()
	serverMux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	serverMux.HandleFunc("/healthz", handlerHealthz)
	server := http.Server{
		Addr:    port,
		Handler: serverMux,
	}

	log.Printf("server started at http://localhost:%s, serving files from :%s\n", port, filepathRoot)
	log.Fatal(server.ListenAndServe())
}
