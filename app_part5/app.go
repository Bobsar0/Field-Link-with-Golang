//PART5:IMPLEMENTING PARSE FILE: TO parse content of several pages and combine into just one page . Displays only one page

package main

import (
	"html/template"
	"io"
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
	r.HandleFunc("/hello", HelloServer)
	r.HandleFunc("/auth", AuthServer)
	//r.HandleFunc("/ik", iksitehandler)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

	log.Fatalln(http.ListenAndServe(":8082", nil))

} //end of main() function

//handler func(ResponseWriter, *Request)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	//func ServeFile(w ResponseWriter, r *Request, name string)
	//http.ServeFile(w, req, "./template/index.html") //the dot indicates to go to the specified location from the root
	// func ParseFiles(filenames ...string) (*Template, error)
	//tmpl, err := template.New("index.html").ParseFiles("index.html", "header.html", "nav.html", "footer.html")//ParseFiles parses like two files
	tmpl := template.Must(template.New("index.html").ParseFiles("./templates/index.html", "./templates/header.html", "./templates/footer.html", "./templates/nav.html")) //ParseGlob parses a whole website

	//auth := false
	//func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	err := tmpl.ExecuteTemplate(w, "index", nil) //execute template handles as a html file
	if err != nil {
		panic(err)
	}

	// func (t *Template) Execute(wr io.Writer, data interface{}) error: requires a parsed CONTENT of the html file and handles as html config/content

}

func AuthServer(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "Now in the Auth Handle func to be Authenticated\n")

}

//using the servefile function

// func ParseFiles(filenames ...string) (*Template, error)
//     ParseFiles creates a new Template and parses the template definitions from
//     the named files. The returned template's name will have the (base) name and
//     (parsed) contents of the first file. There must be at least one file. If an
//     error occurs, parsing stops and the returned *Template is nil.

//     When parsing multiple files with the same name in different directories, the
//     last one mentioned will be the one that results. For instance,
//     ParseFiles("a/foo", "b/foo") stores "b/foo" as the template named "foo",
//     while "a/foo" is unavailable.

//EXECUTE
// func (t *Template) Execute(wr io.Writer, data interface{}) error
//     Execute applies a parsed template to the specified data object, writing the
//     output to wr. If an error occurs executing the template or writing its
//     output, execution stops, but partial results may already have been written
//     to the output writer. A template may be executed safely in parallel.
