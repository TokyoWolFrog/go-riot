package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type routerStruct struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}
type routerList []routerStruct

func allRouters() routerList {
	allRouters := routerList{
		routerStruct{"Index", "GET", "/", indexHandler},
		routerStruct{"Todo", "GET", "/v1/todo", todoHandler},
		routerStruct{"Demo", "GET", "/v1/form", formHandler},
		routerStruct{"Demo", "GET", "/v1/demo", demoHandler},
		routerStruct{"GetTest", "GET", "/v1/user/:user", userHandler},
		routerStruct{"PostTest", "POST", "/v1/post/json", postJSONHandler},
	}
	return allRouters
}

func newRouter() *httprouter.Router {
	newRouter := httprouter.New()

	allRouters := allRouters()
	for _, router := range allRouters {
		var handle httprouter.Handle
		handle = router.HandlerFunc
		handle = handleLogger(handle)

		newRouter.Handle(router.Method, router.Path, handle)
	}

	newRouter.ServeFiles("/static/*filepath", http.Dir("static"))

	return newRouter
}

func handleLogger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		fn(w, r, param)
		log.Printf("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}
