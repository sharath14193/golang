package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /form", formHandler)

	err := http.ListenAndServe(":4001", mux)
	if err != nil {
		log.Fatal("Error occurred while starting the server:", err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data here
  r.ParseForm()
  // Retrieve form value here
  blogName := r.FormValue("blogName")
  blogContent := r.FormValue("blogContent")
	log.Printf("Blog Name: %s\n", blogName)
	log.Printf("Blog Content: %s\n", blogContent)
}
