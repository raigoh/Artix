package functions

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shutting down server...")

	// Set HTTP status code to OK (200)
	w.WriteHeader(http.StatusOK)

	// Define empty data for rendering
	data := IndexPageData{}

	// Render the shutdown confirmation HTML with empty data
	renderTemplate(w, "templates/shutdown.html", data)

	// Gracefully shutdown the server after a delay
	go func() {
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}
