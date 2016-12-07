//PART6:IMPLEMENTING PARSEGLOB and getting two base files to collabo/display

package main

import (
	"html/template"
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
	r.HandleFunc("/index", IndexServer)
	r.HandleFunc("/shop", ShopServer)
	//r.HandleFunc("/auth", AuthServer)
	//r.HandleFunc("/ik", iksitehandler)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

	log.Fatalln(http.ListenAndServe(":8082", nil))

} //end of main() function

//handler func(ResponseWriter, *Request)

// type appTemplate struct {
// 	t *template.Template
// }
type appTemplateInt interface {
	ExecuteT()
}

func IndexServer(w http.ResponseWriter, req *http.Request) {

	var tmpls appTemplateInt
	//func ServeFile(w ResponseWriter, r *Request, name string)
	//http.ServeFile(w, req, "./template/index.html") //the dot indicates to go to the specified location from the root
	// func ParseFiles(filenames ...string) (*Template, error)
	//tmpl, err := template.New("index.html").ParseFiles("index.html", "header.html", "nav.html", "footer.html")
	tmpls.t = template.Must(template.New("index.html").ParseGlob("./templates/*.html"))

	//auth := false
	//func (t *template) Execute(w http.ResponseWriter, req *http.Request, view string, data interface{})

	tmpls.ExecuteT(w, req, "index", nil)
	// futmplnc (t *Template) Execute(wr io.Writer, data interface{}) error: requires a parsed CONTENT of the html file and handles as html config/content
}
func ShopServer(w http.ResponseWriter, req *http.Request) {

	var tmpls appTemplate
	// func ParseFiles(filenames ...string) (*Template, error)
	//tmpl, err := template.New("index.html").ParseFiles("index.html", "header.html", "nav.html", "footer.html")
	tmpls.t = template.Must(template.New("index.html").ParseGlob("./templates/*.html"))

	//auth := false
	//func (t *template) Execute(w http.ResponseWriter, req *http.Request, view string, data interface{})

	tmpls.ExecuteT(w, req, "shop", nil)
	// futmplnc (t *Template) Execute(wr io.Writer, data interface{}) error: requires a parsed CONTENT of the html file and handles as html config/content
}

func (tmpl appTemplate) ExecuteT(w http.ResponseWriter, req *http.Request, view string, data interface{}) {
	err := tmpl.t.ExecuteTemplate(w, view, data) //execute template handles as a html file
	if err != nil {
		panic(err)
	}

}

// func AuthServer(w http.ResponseWriter, req *http.Request) {

// 	tmpl := template.Must(template.New("shop.html").ParseGlob("./templates/*.html"))

// 	//auth := false
// 	//func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
// 	err := tmpl.ExecuteTemplate(w, "shop", nil) //execute template handles as a html file
// 	if err != nil {
// 		panic(err)
// 	}
// }

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

// //ExecuteTemplate
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
//     ExecuteTemplate applies the template associated with t that has the given
//     name to the specified data object and writes the output to wr. If an error
//     occurs executing the template or writing its output, execution stops, but
//     partial results may already have been written to the output writer. A
//     template may be executed safely in parallel.
