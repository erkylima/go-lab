package webapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/erkylima/golab/pkg/route"
	"github.com/gorilla/mux"
)

func ListenAndServe() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	router.HandleFunc("/posts", route.GetPosts).Methods("GET")
	router.HandleFunc("/posts", route.AddPost).Methods("POST")
	log.Println("Server started on port 8000")
	log.Fatalln(http.ListenAndServe(port, router))
}
