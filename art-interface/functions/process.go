package functions

import (
	artdec "art/art-decode"
	"net/http"
)

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate content type for the response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		// If parsing fails, return a bad request error
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Retrieve inputText and operation parameters from the form data
	inputText := r.FormValue("inputText")
	operation := r.FormValue("operation")

	// Check if operation is either decode or encode
	if operation != "decode" && operation != "encode" {
		// If operation is neither decode nor encode, return a bad request error
		http.Error(w, "Malformed query string", http.StatusBadRequest)
		return
	}

	// Define variables for result and status code
	var result string
	var status int

	// Perform the requested operation
	switch operation {
	case "decode":
		result, err = artdec.DecodeMultiLine(inputText)
	case "encode":
		result, err = artdec.EncodeMultiLine(inputText)
	}

	// Check if an error occurred during the operation
	if err != nil {
		// If an error occurred, set status code to 400
		status = http.StatusBadRequest
		result = "Malformed query string"
	} else {
		// If operation was successful, set status code to 202
		status = http.StatusAccepted
	}

	// Define the data to be sent to the template
	data := IndexPageData{
		Result:     result,
		IsDecoding: operation == "decode",
		IsEncoding: operation == "encode",
		HTTPStatus: status,
	}

	// Render the template with appropriate data and status code
	renderTemplate(w, "templates/index.html", data)
}
