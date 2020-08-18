package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"server/articles"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mockArticles := articles.GetMockArticles()
	main := filepath.Join("public", "html", "articles.html")
	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = tmpl.ExecuteTemplate(rw, "articles", mockArticles)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}