package main

import (
	"fmt"
	"log"
	"net/http"

	"go-network-ui/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/submit", handlers.SubmitHandler)
	// http.HandleFunc("/results", handlers.ResultsHandler)
	http.HandleFunc("/results", handlers.ResultsHandler)

	http.HandleFunc("/upload", handlers.UploadHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/download", handlers.DownloadHandler)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
