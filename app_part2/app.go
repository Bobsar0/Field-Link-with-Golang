package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	//1)listen on network net on address addr (listener)
	//2)accept connection (conn)
	//3)read from conn and get the ResponseWriter and the Request
	//4)Route the request/writer to handler
	//5)The handler prepares the response and write it to the writer

	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//HandleFunc registers the handler function for the given pattern in the
	//DefaultServeMux. The documentation for ServeMux explains how patterns are
	//matched.
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/auth", AuthServer)

	//func ListenAndServe(addr string, handler Handler) error
	//ListenAndServe listens on the TCP network address addr and then calls Serve
	//with handler to handle requests on incoming connections. Accepted
	//connections are configured to enable TCP keep-alives. Handler is typically
	//nil, in which case the DefaultServeMux is used.

	//Package log implements a simple logging package. It defines a type, Logger,
	//with methods like func Fatalln(v ...interface{}) for formatting output.
	log.Fatalln(http.ListenAndServe(":8001", nil))

} //end of main() function

//handler func(ResponseWriter, *Request)
func HelloServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "<h1>Biz9ja</h1>")
	io.WriteString(w, "<p>We shall get it's a matter of time</p>")

}

func AuthServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Now in the Auth Handle func to be Authenticated\n")

}
