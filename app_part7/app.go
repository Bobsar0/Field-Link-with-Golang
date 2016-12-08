//PART6:IMPLEMENTING PARSEGLOB and getting two base files to collabo/display
//part7: move PARSEGLOB to another function and call the functon reloadTemplate. Then host a directory for css and img.
//also checkout the dot: if...end

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
	r := mux.NewRouter()
	r.HandleFunc("/index", IndexServer)
	r.HandleFunc("/shop", ShopServer)
	//r.HandleFunc("/auth", AuthServer)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

	log.Fatalln(http.ListenAndServe(":8082", nil))

} //end main()

// type appTemplateInt interface{
// 	ExecuteTemplate()
// }
type appTemplate struct {
	t *template.Template
}

func IndexServer(w http.ResponseWriter, req *http.Request) {

	var tmpls appTemplate
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
	tmpls.t = template.Must(template.New("index.html").ParseGlob("./templates/*.html"))
	tmpls.ExecuteT(w, req, "shop", nil)
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

// EXECUTE
// func (t *Template) Execute(wr io.Writer, data interface{}) error
//     Execute applies a parsed template to the specified data object, writing the
//     output to wr. If an error occurs executing the template or writing its
//     output, execution stops, but partial results may already have been written
//     to the output writer. A template may be executed safely in parallel.

// ExecuteTemplate
// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
//     ExecuteTemplate applies the template associated with t that has the given
//     name to the specified data object and writes the output to wr. If an error
//     occurs executing the template or writing its output, execution stops, but
//     partial results may already have been written to the output writer. A
//     template may be executed safely in parallel.

// MUST (from html/template)
// func Must(t *Template, err error) *Template
//     Must is a helper that wraps a call to a function returning (*Template,
//     error) and panics if the error is non-nil. It is intended for use in
//     variable initializations such as

//     var t = template.Must(template.New("name").Parse("html"))

// NEW (from html/template)
// func New(name string) *Template
//     New allocates a new HTML template with the given name.
