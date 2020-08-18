package controller

import (
	"github.com/denchick/news/database"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func StartPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	mockArticles := database.GetMockArticles()
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