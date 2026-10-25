package tmdb

type AuthenticationResponse struct {
	Success       bool   `json:"success"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

type MovieSearchResponse struct {
	Page         int                `json:"page"`
	Results      []MovieSearchResult `json:"results"`
	TotalPages   int                `json:"total_pages"`
	TotalResults int                `json:"total_results"`
}

type MovieSearchResult struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	PosterPath       *string `json:"poster_path"`
	BackdropPath     *string `json:"backdrop_path"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Popularity       float64 `json:"popularity"`
	Adult            bool    `json:"adult"`
	OriginalLanguage string  `json:"original_language"`
}

type MovieDetails struct {
	ID               int              `json:"id"`
	Title            string           `json:"title"`
	OriginalTitle    string           `json:"original_title"`
	Overview         string           `json:"overview"`
	Tagline          string           `json:"tagline"`
	ReleaseDate      string           `json:"release_date"`
	Runtime          int              `json:"runtime"`
	Status           string           `json:"status"`
	PosterPath       *string          `json:"poster_path"`
	BackdropPath     *string          `json:"backdrop_path"`
	VoteAverage      float64          `json:"vote_average"`
	VoteCount        int              `json:"vote_count"`
	Popularity       float64          `json:"popularity"`
	Budget           int              `json:"budget"`
	Revenue          int              `json:"revenue"`
	Adult            bool             `json:"adult"`
	OriginalLanguage string           `json:"original_language"`
	Homepage         string           `json:"homepage"`
	IMDBID           string           `json:"imdb_id"`
	Genres           []NamedEntity    `json:"genres"`
	ProductionCompanies []NamedEntity `json:"production_companies"`
}

type NamedEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type tmdbErrorResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}
