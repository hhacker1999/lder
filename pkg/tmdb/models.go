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
	Genres              []NamedEntity `json:"genres"`
	ProductionCompanies []NamedEntity `json:"production_companies"`
	Credits             *MediaCredits `json:"credits"`
}

type MediaCredits struct {
	Cast []CastMember `json:"cast"`
	Crew []CrewMember `json:"crew"`
}

type CastMember struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Character   string  `json:"character"`
	Order       int     `json:"order"`
	ProfilePath *string `json:"profile_path"`
}

type CrewMember struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Job         string  `json:"job"`
	Department  string  `json:"department"`
	ProfilePath *string `json:"profile_path"`
}

type NamedEntity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TVSearchResponse struct {
	Page         int              `json:"page"`
	Results      []TVSearchResult `json:"results"`
	TotalPages   int              `json:"total_pages"`
	TotalResults int              `json:"total_results"`
}

type TVSearchResult struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	FirstAirDate     string  `json:"first_air_date"`
	PosterPath       *string `json:"poster_path"`
	BackdropPath     *string `json:"backdrop_path"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Popularity       float64 `json:"popularity"`
	OriginalLanguage string  `json:"original_language"`
}

type TVExternalIDs struct {
	ID     int    `json:"id"`
	IMDBID string `json:"imdb_id"`
	TVDBID int    `json:"tvdb_id"`
}

type TVDetails struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	OriginalName      string        `json:"original_name"`
	Overview          string        `json:"overview"`
	PosterPath        *string       `json:"poster_path"`
	BackdropPath      *string       `json:"backdrop_path"`
	FirstAirDate      string        `json:"first_air_date"`
	LastAirDate       string        `json:"last_air_date"`
	Status            string        `json:"status"`
	Type              string        `json:"type"`
	NumberOfSeasons   int           `json:"number_of_seasons"`
	NumberOfEpisodes  int           `json:"number_of_episodes"`
	EpisodeRunTime    []int         `json:"episode_run_time"`
	VoteAverage       float64       `json:"vote_average"`
	VoteCount         int           `json:"vote_count"`
	OriginalLanguage  string        `json:"original_language"`
	Genres            []NamedEntity `json:"genres"`
	Networks          []NamedEntity `json:"networks"`
	CreatedBy         []NamedEntity `json:"created_by"`
	Credits           *MediaCredits `json:"credits"`
}

type tmdbErrorResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}
