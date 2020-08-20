package controller

import (
	"github.com/denchick/news/database"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
)

func HandleIndexPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	articles := database.ReadLastArticles(Db)
	templatePath := filepath.Join("public", "html", "articles.html")
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = template.ExecuteTemplate(rw, "articles", articles)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

func HandleSearchPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	query := r.URL.Query().Get("query")
	articles := database.FindArticleByTitle(Db, query)
	templatePath := filepath.Join("public", "html", "articles.html")
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = template.ExecuteTemplate(rw, "articles", articles)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}