package main

import (
	"net/http"
	"os"
)

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthzResponse{
		Response: Response{
			Code:    http.StatusOK,
			Message: "success",
		},
	}
	ServiceLog(r.Host, http.StatusOK)
	WriteStruct(w, resp)
}

func ReqHeaderHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
	resp := Response{
		Code:    0,
		Message: "success",
	}
	ServiceLog(r.Host, http.StatusOK)
	WriteStruct(w, resp)
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	resp := VersionResponse{
		Response: Response{
			Code:    0,
			Message: "success",
		},
		Version: os.Getenv("VERSION"),
	}
	ServiceLog(r.Host, http.StatusOK)
	WriteStruct(w, resp)
}
