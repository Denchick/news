package controller

import (
	"fmt"
	"github.com/denchick/news/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleSearchPage(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	query := r.URL.Query().Get("query")
	result := database.FindArticleByTitle(Db, query)
	fmt.Fprint(rw, result)
}