package router

import (
	"lder/pkg/tmdb"
	"lder/pkg/tvdb"
	"strings"
)

type searchResponse struct {
	Movies     []mediaSearchItem `json:"movies"`
	Shows      []mediaSearchItem `json:"shows"`
	Page       int               `json:"page"`
	TotalPages int               `json:"totalPages"`
}

type mediaSearchItem struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	Year        string  `json:"year"`
	PosterURL   string  `json:"posterUrl"`
	VoteAverage float64 `json:"voteAverage"`
	MediaType   string  `json:"mediaType"`
}

type castMemberDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	ProfileURL  string `json:"profileUrl"`
}

type crewMemberDTO struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

type movieDetailResponse struct {
	ID                  int             `json:"id"`
	Title               string          `json:"title"`
	OriginalTitle       string          `json:"originalTitle"`
	Overview            string          `json:"overview"`
	Tagline             string          `json:"tagline"`
	ReleaseDate         string          `json:"releaseDate"`
	Runtime             int             `json:"runtime"`
	Status              string          `json:"status"`
	PosterURL           string          `json:"posterUrl"`
	BackdropURL         string          `json:"backdropUrl"`
	VoteAverage         float64         `json:"voteAverage"`
	VoteCount           int             `json:"voteCount"`
	Genres              []string        `json:"genres"`
	IMDBID              string          `json:"imdbId"`
	Homepage            string          `json:"homepage"`
	OriginalLanguage    string          `json:"originalLanguage"`
	Budget              int             `json:"budget"`
	Revenue             int             `json:"revenue"`
	ProductionCompanies []string        `json:"productionCompanies"`
	Directors           []string        `json:"directors"`
	Cast                []castMemberDTO `json:"cast"`
}

type showDetailResponse struct {
	ID                 int             `json:"id"`
	TmdbID             int             `json:"tmdbId"`
	TvdbID             int             `json:"tvdbId"`
	Name               string          `json:"name"`
	OriginalName       string          `json:"originalName"`
	Overview           string          `json:"overview"`
	Status             string          `json:"status"`
	TmdbStatus         string          `json:"tmdbStatus"`
	FirstAired         string          `json:"firstAired"`
	LastAired          string          `json:"lastAired"`
	NextAired          string          `json:"nextAired"`
	Year               string          `json:"year"`
	Score              float64         `json:"score"`
	TmdbVoteAverage    float64         `json:"tmdbVoteAverage"`
	TmdbVoteCount      int             `json:"tmdbVoteCount"`
	ImageURL           string          `json:"imageUrl"`
	BackdropURL        string          `json:"backdropUrl"`
	AverageRuntime     *int            `json:"averageRuntime"`
	EpisodeRunTime     []int           `json:"episodeRunTime"`
	NumberOfSeasons    int             `json:"numberOfSeasons"`
	NumberOfEpisodes   int             `json:"numberOfEpisodes"`
	OriginalCountry    string          `json:"originalCountry"`
	OriginalLanguage   string          `json:"originalLanguage"`
	IMDBID             string          `json:"imdbId"`
	Genres             []string        `json:"genres"`
	Networks           []string        `json:"networks"`
	CreatedBy          []string        `json:"createdBy"`
	LastUpdated        string          `json:"lastUpdated"`
	ShowType           string          `json:"showType"`
	Cast               []castMemberDTO `json:"cast"`
}

func yearFromDate(date string) string {
	if date == "" {
		return ""
	}
	parts := strings.Split(date, "-")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func mapMovieSearchResults(results []tmdb.MovieSearchResult) []mediaSearchItem {
	items := make([]mediaSearchItem, 0, len(results))
	for _, m := range results {
		items = append(items, mediaSearchItem{
			ID:          m.ID,
			Title:       m.Title,
			Overview:    m.Overview,
			Year:        yearFromDate(m.ReleaseDate),
			PosterURL:   tmdb.PosterURL(m.PosterPath),
			VoteAverage: m.VoteAverage,
			MediaType:   "movie",
		})
	}
	return items
}

func mapTVSearchResults(results []tmdb.TVSearchResult) []mediaSearchItem {
	items := make([]mediaSearchItem, 0, len(results))
	for _, s := range results {
		items = append(items, mediaSearchItem{
			ID:          s.ID,
			Title:       s.Name,
			Overview:    s.Overview,
			Year:        yearFromDate(s.FirstAirDate),
			PosterURL:   tmdb.PosterURL(s.PosterPath),
			VoteAverage: s.VoteAverage,
			MediaType:   "show",
		})
	}
	return items
}

func mapCast(credits *tmdb.MediaCredits, limit int) []castMemberDTO {
	if credits == nil {
		return nil
	}
	cast := credits.Cast
	if len(cast) > limit {
		cast = cast[:limit]
	}
	out := make([]castMemberDTO, 0, len(cast))
	for _, c := range cast {
		out = append(out, castMemberDTO{
			ID:         c.ID,
			Name:       c.Name,
			Role:       c.Character,
			ProfileURL: tmdb.PosterURL(c.ProfilePath),
		})
	}
	return out
}

func mapDirectors(credits *tmdb.MediaCredits) []string {
	if credits == nil {
		return nil
	}
	var directors []string
	for _, c := range credits.Crew {
		if c.Job == "Director" {
			directors = append(directors, c.Name)
		}
	}
	return directors
}

func mapMovieDetails(m tmdb.MovieDetails) movieDetailResponse {
	genres := make([]string, 0, len(m.Genres))
	for _, g := range m.Genres {
		genres = append(genres, g.Name)
	}
	companies := make([]string, 0, len(m.ProductionCompanies))
	for _, c := range m.ProductionCompanies {
		companies = append(companies, c.Name)
	}
	return movieDetailResponse{
		ID:                  m.ID,
		Title:               m.Title,
		OriginalTitle:       m.OriginalTitle,
		Overview:            m.Overview,
		Tagline:             m.Tagline,
		ReleaseDate:         m.ReleaseDate,
		Runtime:             m.Runtime,
		Status:              m.Status,
		PosterURL:           tmdb.PosterURL(m.PosterPath),
		BackdropURL:         tmdb.BackdropURL(m.BackdropPath),
		VoteAverage:         m.VoteAverage,
		VoteCount:           m.VoteCount,
		Genres:              genres,
		IMDBID:              m.IMDBID,
		Homepage:            m.Homepage,
		OriginalLanguage:    m.OriginalLanguage,
		Budget:              m.Budget,
		Revenue:             m.Revenue,
		ProductionCompanies: companies,
		Directors:           mapDirectors(m.Credits),
		Cast:                mapCast(m.Credits, 24),
	}
}

func mapShowDetails(
	tmdbID int,
	tvdbID int,
	imdbID string,
	s tvdb.SeriesDetails,
	tv tmdb.TVDetails,
) showDetailResponse {
	status := ""
	if s.Status != nil {
		status = s.Status.Name
	}
	imageURL := s.Image
	backdropURL := tmdb.BackdropURL(tv.BackdropPath)
	if backdropURL == "" {
		backdropURL = tmdb.PosterURL(tv.PosterPath)
	}
	if imageURL == "" {
		imageURL = tmdb.PosterURL(tv.PosterPath)
	}

	genres := make([]string, 0, len(tv.Genres))
	for _, g := range tv.Genres {
		genres = append(genres, g.Name)
	}
	networks := make([]string, 0, len(tv.Networks))
	for _, n := range tv.Networks {
		networks = append(networks, n.Name)
	}
	createdBy := make([]string, 0, len(tv.CreatedBy))
	for _, c := range tv.CreatedBy {
		createdBy = append(createdBy, c.Name)
	}

	return showDetailResponse{
		ID:               s.ID,
		TmdbID:           tmdbID,
		TvdbID:           tvdbID,
		Name:             s.Name,
		OriginalName:     tv.OriginalName,
		Overview:         s.Overview,
		Status:           status,
		TmdbStatus:       tv.Status,
		FirstAired:       s.FirstAired,
		LastAired:        s.LastAired,
		NextAired:        s.NextAired,
		Year:             s.Year,
		Score:            s.Score,
		TmdbVoteAverage:  tv.VoteAverage,
		TmdbVoteCount:    tv.VoteCount,
		ImageURL:         imageURL,
		BackdropURL:      backdropURL,
		AverageRuntime:   s.AverageRuntime,
		EpisodeRunTime:   tv.EpisodeRunTime,
		NumberOfSeasons:  tv.NumberOfSeasons,
		NumberOfEpisodes: tv.NumberOfEpisodes,
		OriginalCountry:  s.OriginalCountry,
		OriginalLanguage: s.OriginalLanguage,
		IMDBID:           imdbID,
		Genres:           genres,
		Networks:         networks,
		CreatedBy:        createdBy,
		LastUpdated:      s.LastUpdated,
		ShowType:         tv.Type,
		Cast:             mapCast(tv.Credits, 24),
	}
}
