package main

import (
	"log"
	"net/http"

	"github.com/denchick/news/server/controller"
	"github.com/julienschmidt/httprouter"
)

func createRouter() (router *httprouter.Router) {
	router = httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))
	router.GET("/", controller.HandleIndexPage)
	router.GET("/search", controller.HandleSearchPage)
	return
}

func main() {
	err := controller.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	addr := ":8081"
	log.Printf("run server on %s", addr)
	err = http.ListenAndServe(addr, createRouter())
	if err != nil {
		log.Fatal(err)
	}
}
