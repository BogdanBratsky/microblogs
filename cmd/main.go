package main

import (
	"log"
	"net/http"

	"github.com/BogdanBratsky/microblogs/internal/auth"
)

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err.Error())
	}

	http.HandleFunc("/api/v1/auth/register", auth.RegisterHandler)
}
