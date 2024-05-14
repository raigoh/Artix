package functions

import (
	artdec "art/art-decode"
	"net/http"
)

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	inputText := r.FormValue("inputText")
	isDecoding := r.FormValue("operation") == "decode"
	isEncoding := r.FormValue("operation") == "encode"

	var result string
	var success bool
	if isDecoding {
		result, err = artdec.DecodeMultiLine(inputText)
		if err != nil {
			success = false
		} else {
			success = true
		}
	} else if isEncoding {
		result, err = artdec.EncodeMultiLine(inputText)
		if err != nil {
			success = false
		} else {
			success = true
		}
	}

	data := IndexPageData{
		Result:     result,
		IsDecoding: isDecoding,
		IsEncoding: isEncoding,
	}

	if !success {
		data.HTTPStatus = http.StatusBadRequest
		data.Result = "Malformed query string."
		renderTemplate(w, "templates/index.html", data)
		return
	}

	data.HTTPStatus = http.StatusAccepted
	renderTemplate(w, "templates/index.html", data)
}
