package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var (
	todoTplPath = "html/todo_tpl.html"
	formTplPath = "html/form_tpl.html"
	demoTplPath = "html/demo_tpl.html"
)

type demoStruct struct {
	Title   string
	Detail1 string
	Detail2 string
	Detail3 string
	Detail4 string
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allRouters := allRouters()
	var pathList []string
	for _, router := range allRouters {
		pathList = append(pathList, router.Path)
	}
	fmt.Fprintf(w, "Hello! valid url is: "+"\r"+strings.Join(pathList, "\r"))
}

func todoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles(todoTplPath))
	if err := t.Execute(w, todoTplPath); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles(formTplPath))
	if err := t.Execute(w, formTplPath); err != nil {
		log.Fatal(err)
	}
}

func demoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	demoList := []demoStruct{
		{"title1", "t1-detail1", "t1-detail2", "t1-detail3", "t1-detail4"},
		{"title2", "t2-detail1", "t2-detail2", "t2-detail3", "t2-detail4"},
		{"title3", "t3-detail1", "t3-detail2", "t3-detail3", "t3-detail4"},
		{"title4", "t4-detail1", "t4-detail2", "t4-detail3", "t4-detail4"},
		{"title5", "t5-detail1", "t5-detail2", "t5-detail3", "t5-detail4"},
		{"title6", "t6-detail1", "t6-detail2", "t6-detail3", "t6-detail4"},
		{"A", "detail1", "detail2", "detail3", "detail4"},
		{"a", "detail1", "detail2", "detail3", "detail4"},
		{"B", "detail1", "detail2", "detail3", "detail4"},
		{"C", "detail1", "detail2", "detail3", "detail4"},
		{"ab", "detail1", "detail2", "detail3", "detail4"},
		{"Abc", "detail1", "detail2", "detail3", "detail4"},
		{"日本語テスト", "detail1", "detail2", "detail3", "detail4"},
		{"何か良い", "detail1", "detail2", "detail3", "detail4"},
		{"日曜日", "detail1", "detail2", "detail3", "detail4"},
	}
	jsonBytes, _ := json.Marshal(demoList)

	demoTplName := path.Base(demoTplPath)
	t := template.Must(template.ParseFiles(demoTplPath))
	if err := t.ExecuteTemplate(w, demoTplName, string(jsonBytes)); err != nil {
		log.Fatal(err)
	}
}

type userStruct struct {
	User string
}

func userHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userJSON := userStruct{"Hello, " + ps.ByName("user") + "!(from Golang server)"}
	jsonBytes, _ := json.Marshal(userJSON)
	fmt.Fprintf(w, string(jsonBytes))
}

type formStruct struct {
	Select   string
	Name     string
	Password string
	Email    string
	Foo      string
	Cb       string
}

func postJSONHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var form formStruct
	json.NewDecoder(r.Body).Decode(&form)

	jsonBytes, _ := json.Marshal(form)
	fmt.Fprintf(w, string(jsonBytes))
}
