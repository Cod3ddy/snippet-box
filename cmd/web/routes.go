package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//using the mux.Handle() to register the file server as the hande for all the URL paths that start with "/static/" . For matching paths, we strip the "/static" before the request reaches the file server.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.SessionManger.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost))

	// middleware chain containt 'standard' middleware.
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	//Wrap the existing chain with the logRequest middleware
	return standard.Then(mux)
}
