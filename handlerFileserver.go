package main

import "net/http"

func handlerAssets(prefix, filepathRoot string) http.Handler {
	return http.StripPrefix(prefix, http.FileServer(http.Dir(filepathRoot)))
}
