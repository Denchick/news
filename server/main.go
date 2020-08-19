package main

import (
	"github.com/denchick/news/server/controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func createRouter() (router *httprouter.Router){
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
	err = http.ListenAndServe("localhost:4444", createRouter())
	if err != nil {
		log.Fatal(err)
	}
}