package functions

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shutting down server...")
	// Render the shutdown confirmation HTML
	renderTemplate(w, "templates/shutdown.html", nil)
	// Gracefully shutdown the server after a delay
	go func() {
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
}
