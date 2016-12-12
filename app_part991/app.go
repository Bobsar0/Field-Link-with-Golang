package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Use gorilla/mux for rich routing.
	// See http://www.gorillatoolkit.org/pkg/mux
	r := mux.NewRouter()
	r.HandleFunc("/", firstHandler)

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	// [END request_logging]

	log.Fatalln(http.ListenAndServe(":8000", nil))

}

func firstHandler(w http.ResponseWriter, r *http.Request) {
	//var firstTmpl *appTemplate

	firstTmpl := parseTemplate("first.html")

	if err := firstTmpl.Execute(w, r, nil); err != nil {
		panic(err)

	}
}

type appError struct {
	Error   error
	Message string
	Code    int
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}

