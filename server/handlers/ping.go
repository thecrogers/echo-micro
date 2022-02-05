package handlers

import (
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("OK"))
}
