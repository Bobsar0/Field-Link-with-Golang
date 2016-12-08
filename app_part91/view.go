package main

import (
	"html/template"
	"net/http"
)

type appTemplate struct {
	t    *template.Template //final page template
	page []string //name of templates that form t used by parsefiles
}

//("./templates/index.html", "./templates/header.html", "./templates/footer.html", "./templates/nav.html"))
func parseTemplate(ppage []string) *appTemplate {
	var tmpl appTemplate
	tmpl.page = ppage
	tmpl.t = template.Must(template.New("index.html").ParseFiles(tmpl.page...))
	return &tmpl
}
func (tmpl *appTemplate) ExecuteT(w http.ResponseWriter, req *http.Request, view string, data interface{}) {
	err := tmpl.t.ExecuteTemplate(w, view, data) //execute template handles as a html file
	if err != nil {
		panic(err)
	}
}
