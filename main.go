package main

import (
	"log"
	"net/http"
)

func main() {
	filepathRoot := "."
	port := ":8080"
	apiCfg := apiConfig{}
	serverMux := http.NewServeMux()

	serverMux.Handle("/app/", apiCfg.middlewareMetricsInc(handlerAssets("/app", filepathRoot)))
	serverMux.HandleFunc("/healthz", handlerReadiness)
	serverMux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	serverMux.HandleFunc("/reset", apiCfg.handlerResetMetrics)

	server := http.Server{
		Addr:    port,
		Handler: serverMux,
	}

	log.Printf("server started at http://localhost:%s, serving files from :%s\n", port, filepathRoot)
	log.Fatal(server.ListenAndServe())
}
