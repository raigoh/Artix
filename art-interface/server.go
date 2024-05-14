package main

import (
	"art/art-interface/functions"
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.HandleFunc("/", functions.IndexHandler)
	http.HandleFunc("/process", functions.ProcessHandler)
	http.HandleFunc("/shutdown", functions.ShutdownHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
