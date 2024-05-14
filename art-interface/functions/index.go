package functions

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	data := IndexPageData{
		HTTPStatus: http.StatusOK,
		Result:     "",
	}
	renderTemplate(w, "templates/index.html", data)
}
