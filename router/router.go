package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HttpRouter struct {
	router *chi.Mux
}

func New() HttpRouter {
	return HttpRouter{}
}

func (r *HttpRouter) SetupRoutes() http.Handler {
	r.router = chi.NewRouter()
	r.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Status Ok"))
	})

	return r.router
}
