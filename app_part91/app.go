//PART6:IMPLEMENTING PARSEGLOB and getting two base files to collabo/display
//part7: move PARSEGLOB to another function and call the functon reloadTemplate. Then host a directory for css and img.
//also checkout the dot: if...end

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var indexpage = []string{"./templates/index.html", "./templates/header.html", "./templates/footer.html", "./templates/nav.html"}
var shoppage = []string{"./templates/shop.html", "./templates/header2.html", "./templates/footer.html"}
var indextmpl = parseTemplate(indexpage)
var shoptmpl = parseTemplate(shoppage)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexServer)
	r.HandleFunc("/shop", ShopServer)
	r.PathPrefix("/asset/static/").
		Handler(http.StripPrefix("/asset/static/", http.FileServer(http.Dir("./asset/static"))))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	log.Fatalln(http.ListenAndServe(":8082", nil))

} //end main()
func IndexServer(w http.ResponseWriter, req *http.Request) {
	indextmpl.ExecuteT(w, req, "index", nil)
}
func ShopServer(w http.ResponseWriter, req *http.Request) {
	shoptmpl.ExecuteT(w, req, "shop", nil)
}

