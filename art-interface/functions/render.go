package functions

import (
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the HTTP status code in the response
	status := http.StatusOK
	if d, ok := data.(IndexPageData); ok {
		status = d.HTTPStatus
	}
	w.WriteHeader(status)

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
