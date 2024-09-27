package users

import (
	"log"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("request recieved", r.Method, r.URL)
}
