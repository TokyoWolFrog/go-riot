package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	todoTplPath = "html/todo_tpl.html"
)

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello! valid url is: /v1/todo")
}

func todoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles(todoTplPath))
	if err := t.Execute(w, todoTplPath); err != nil {
		log.Fatal(err)
	}
}
