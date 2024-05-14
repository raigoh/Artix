package functions

// IndexPageData represents data for the index page template.
type IndexPageData struct {
	HTTPStatus int    // HTTP status code
	Result     string // Result to be displayed on the page
	IsDecoding bool   // Indicates if decoding functionality is selected
	IsEncoding bool   // Indicates if encoding functionality is selected
}
