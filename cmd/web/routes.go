package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	//Note the path directory given to the http.Dir function is relative to the project directory root.

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//using the mux.Handle() to register the file server as the hande for all the URL paths that start with "/static/" . For matching paths, we strip the "/static" before the request reaches the file server.
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return mux
}
