package main

import (
	"net/http"
	"io"
)

type StringHandler struct {
	message string
}

// The browser makes two HTTP requests.
// The first is for /, which is the path component of the URL that was requested.
// The second request is for /favicon.ico, which the browser sends to get an icon
// to display at the top of the window or tab

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
		if (request.URL.Path == "/favicon.ico") {				// net/url: type URL struct {...}
			Printfln("Request for icon detected - returning 404")
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		Printfln("Request for %v", request.URL.Path)
		io.WriteString(writer, sh.message)
	}

	func main() {
		err := http.ListenAndServe(":5000", StringHandler{message: "Hello-Mello Kissy-Missy"})
		if (err != nil) {
			Printfln("Error: %v", err.Error())
		}
	}

	// You may find that subsequent requests from the browser for 
	// http://localhost:5000 do not trigger a second request for the icon file.
	// That’s because the browser notes the 404 response and knows that there
	// is no icon file for this URL. Clear the browser’s cache and request
	// http://localhost:5000 to return to the original behavior.
