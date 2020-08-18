package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Search(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	text := "Приветствую тебя на стартовой странице этого сайта!"
	//возвращаем простой текст
	fmt.Fprint(rw, text)
}