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
<<<<<<< HEAD
=======
	//r.HandleFunc("/ik", iksitehandler)
>>>>>>> c0da7012638fb768548c71bbb7547beb6173fe5c

	//func ListenAndServe(addr string, handler Handler) error
	//ListenAndServe listens on the TCP network address addr and then calls Serve
	//with handler to handle requests on incoming connections. Accepted
	//connections are configured to enable TCP keep-alives. Handler is typically
	//nil, in which case the DefaultServeMux is used.

	//Package log implements a simple logging package. It defines a type, Logger,
	//with methods like func Fatalln(v ...interface{}) for formatting output.
	log.Fatalln(http.ListenAndServe(":8001", r))

} //end of main() function
<<<<<<< HEAD

//handler func(ResponseWriter, *Request)
func HelloServer(w http.ResponseWriter, req *http.Request) {
	//func ParseFiles(filenames ...string) (*Template, error)
	//tmpl, err := template.ParseFiles()
	//func ServeFile(w ResponseWriter, r *Request, name string)
	http.ServeFile(w, req, "./template/index.html")

}

func AuthServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Now in the Auth Handle func to be Authenticated\n")

}
=======

//handler func(ResponseWriter, *Request)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	//func ServeFile(w ResponseWriter, r *Request, name string)
	http.ServeFile(w, req, "./template/index.html") //the dot indicates to go to the specified location from the root
	// func ParseFiles(filenames ...string) (*Template, error)
}

func AuthServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Now in the Auth Handle func to be Authenticated\n")

}

//using the servrefile function

// func ParseFiles(filenames ...string) (*Template, error)
//     ParseFiles creates a new Template and parses the template definitions from
//     the named files. The returned template's name will have the (base) name and
//     (parsed) contents of the first file. There must be at least one file. If an
//     error occurs, parsing stops and the returned *Template is nil.

//     When parsing multiple files with the same name in different directories, the
//     last one mentioned will be the one that results. For instance,
//     ParseFiles("a/foo", "b/foo") stores "b/foo" as the template named "foo",
//     while "a/foo" is unavailable.
>>>>>>> c0da7012638fb768548c71bbb7547beb6173fe5c
