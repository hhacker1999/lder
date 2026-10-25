package router

import (
	"lder/pkg/tmdb"
	"lder/pkg/tvdb"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpRouter struct {
	router *chi.Mux
	tmdb   tmdb.Tmdb
	tvdb   *tvdb.Tvdb
}

func New(tmdbClient tmdb.Tmdb, tvdbClient *tvdb.Tvdb) HttpRouter {
	return HttpRouter{
		tmdb: tmdbClient,
		tvdb: tvdbClient,
	}
}

func (r *HttpRouter) SetupRoutes() http.Handler {
	r.router = chi.NewRouter()
	r.router.Use(middleware.RequestID)
	r.router.Use(middleware.RealIP)
	r.router.Use(middleware.Recoverer)

	r.router.Route("/api", func(api chi.Router) {
		api.Get("/health", r.handleHealth)
		api.Get("/search", r.handleSearch)
		api.Get("/movies/{tmdbId}", r.handleMovieDetails)
		api.Get("/shows/{tmdbId}", r.handleShowDetails)
	})

	r.mountStatic()

	return r.router
}
