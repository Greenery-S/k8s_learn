package main

import (
	"net/http"
)

func main() {
	// a http server
	http.HandleFunc("/healthz", HealthzHandler)
	http.HandleFunc("/req_header", ReqHeaderHandler)
	http.HandleFunc("/version", VersionHandler)
	http.ListenAndServe("localhost:8080", nil)
}
