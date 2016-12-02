package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Use gorilla/mux for rich routing.
	// See http://www.gorillatoolkit.org/pkg/mux
	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloServer)
	r.HandleFunc("/auth", AuthServer)

	//func ListenAndServe(addr string, handler Handler) error
	//ListenAndServe listens on the TCP network address addr and then calls Serve
	//with handler to handle requests on incoming connections. Accepted
	//connections are configured to enable TCP keep-alives. Handler is typically
	//nil, in which case the DefaultServeMux is used.

	//Package log implements a simple logging package. It defines a type, Logger,
	//with methods like func Fatalln(v ...interface{}) for formatting output.
	log.Fatalln(http.ListenAndServe(":8001", r))

} //end of main() function

//handler func(ResponseWriter, *Request)
func HelloServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "<h1>Biz9ja</h1>")
	io.WriteString(w, "<p>We shall get it's a matter of time</p>")

}

func AuthServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Now in the Auth Handle func to be Authenticated\n")

}
