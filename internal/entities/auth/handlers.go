package auth

import (
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request recieved", r.Method, r.URL)
}
