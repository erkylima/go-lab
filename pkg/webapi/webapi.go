package webapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ListenAndServe() error {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	http.ListenAndServe(port, router)
	log.Println("Server started on port 8000")
}
