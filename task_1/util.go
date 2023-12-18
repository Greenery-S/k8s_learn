package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteStruct(w http.ResponseWriter, stru any) {
	bytes, _ := json.Marshal(stru)
	_, _ = w.Write(bytes)
}

func ServiceLog(ip string, sc int) {
	log.Println(ip, sc)
}
