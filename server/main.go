package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"server/controller"
)

func main() {
	//создаем и запускаем в работу роутер для обслуживания запросов
	r := httprouter.New()
	routes(r)

	//прикрепляемся к хосту и свободному порту для приема и обслуживания входящих запросов
	//вторым параметром передается роутер, который будет работать с запросами
	err := http.ListenAndServe("localhost:4444", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.StartPage)
	r.GET("/search", controller.Search)
}