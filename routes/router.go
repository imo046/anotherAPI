package routes

import (
	"example.com/m/v2/handlers"
	"github.com/gorilla/mux"
)

type ApiRouter struct {
	Router  *mux.Router
	Handler *handlers.Handler
}

func (r *ApiRouter) Home(route string) {
	r.Router.HandleFunc(route, r.Handler.HomeLink)
}

//func (r *ApiRouter) Create(route string) {
//	r.Router.HandleFunc(route, r.Handler.CreateEntry).Methods("POST")
//}
//
//func (r *ApiRouter) CreateMultiple(route string) {
//	r.Router.HandleFunc(route, r.Handler.CreateEntries).Methods("POST")
//}

func (r *ApiRouter) Count(route string) {
	r.Router.HandleFunc(route, r.Handler.CountEntries).Methods("GET")
}

// UPLOAD
func (r *ApiRouter) Upload(route string) {
	r.Router.HandleFunc(route, r.Handler.Upload).Methods("POST")
}

func (r *ApiRouter) GetOne(route string) {
	r.Router.HandleFunc(route, r.Handler.GetEntry).Methods("GET")
}

func (r *ApiRouter) GetAll(route string) {
	r.Router.HandleFunc(route, r.Handler.GetEntries).Methods("GET")
}

func (r *ApiRouter) Delete(route string) {
	r.Router.HandleFunc(route, r.Handler.DeleteOne).Methods("DELETE")
}

func (r *ApiRouter) DeleteAll(route string) {
	r.Router.HandleFunc(route, r.Handler.DeleteAll).Methods("DELETE")
}

func (r *ApiRouter) Update(route string) {
	r.Router.HandleFunc(route, r.Handler.Update).Methods("PATCH")
}
