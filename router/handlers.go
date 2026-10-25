package router

import (
	"encoding/json"
	"errors"
	"lder/pkg/tmdb"
	"lder/pkg/tvdb"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
)

func (r *HttpRouter) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (r *HttpRouter) handleSearch(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("q")
	if query == "" {
		writeError(w, http.StatusBadRequest, "query parameter q is required")
		return
	}

	page := 1
	if p := req.URL.Query().Get("page"); p != "" {
		n, err := strconv.Atoi(p)
		if err != nil || n < 1 {
			writeError(w, http.StatusBadRequest, "invalid page")
			return
		}
		page = n
	}

	var movies tmdb.MovieSearchResponse
	var shows tmdb.TVSearchResponse

	g := new(errgroup.Group)
	g.Go(func() error {
		res, err := r.tmdb.SearchMovies(query, page)
		if err != nil {
			return err
		}
		movies = res
		return nil
	})
	g.Go(func() error {
		res, err := r.tmdb.SearchTV(query, page)
		if err != nil {
			return err
		}
		shows = res
		return nil
	})

	if err := g.Wait(); err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}

	totalPages := movies.TotalPages
	if shows.TotalPages > totalPages {
		totalPages = shows.TotalPages
	}

	writeJSON(w, http.StatusOK, searchResponse{
		Movies:     mapMovieSearchResults(movies.Results),
		Shows:      mapTVSearchResults(shows.Results),
		Page:       page,
		TotalPages: totalPages,
	})
}

func (r *HttpRouter) handleMovieDetails(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDParam(req, "tmdbId")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	details, err := r.tmdb.MovieDetails(id)
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, mapMovieDetails(details))
}

func (r *HttpRouter) handleShowDetails(w http.ResponseWriter, req *http.Request) {
	tmdbID, err := parseIDParam(req, "tmdbId")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	external, err := r.tmdb.TVExternalIDs(tmdbID)
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}
	if external.TVDBID <= 0 {
		writeError(w, http.StatusNotFound, "no TVDB mapping for this series")
		return
	}

	var series tvdb.SeriesDetails
	var tv tmdb.TVDetails

	g := new(errgroup.Group)
	g.Go(func() error {
		res, err := r.tvdb.SeriesDetails(external.TVDBID)
		if err != nil {
			return err
		}
		series = res
		return nil
	})
	g.Go(func() error {
		res, err := r.tmdb.TVDetails(tmdbID)
		if err != nil {
			return err
		}
		tv = res
		return nil
	})
	if err := g.Wait(); err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, mapShowDetails(tmdbID, external.TVDBID, external.IMDBID, series, tv))
}

func parseIDParam(req *http.Request, name string) (int, error) {
	raw := chi.URLParam(req, name)
	if raw == "" {
		return 0, errors.New("missing id")
	}
	id, err := strconv.Atoi(raw)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid id")
	}
	return id, nil
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
