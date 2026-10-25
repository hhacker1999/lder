package tvdb

type loginRequest struct {
	APIKey string `json:"apikey"`
}

type loginResponse struct {
	Status string    `json:"status"`
	Data   loginData `json:"data"`
}

type loginData struct {
	Token string `json:"token"`
}

type SeriesDetailsResponse struct {
	Status string       `json:"status"`
	Data   SeriesDetails `json:"data"`
}

type SeriesDetails struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	Slug              string        `json:"slug"`
	Image             string        `json:"image"`
	Overview          string        `json:"overview"`
	FirstAired        string        `json:"firstAired"`
	LastAired         string        `json:"lastAired"`
	NextAired         string        `json:"nextAired"`
	Year              string        `json:"year"`
	Score             float64       `json:"score"`
	Status            *SeriesStatus `json:"status"`
	OriginalCountry   string        `json:"originalCountry"`
	OriginalLanguage  string        `json:"originalLanguage"`
	AverageRuntime    *int          `json:"averageRuntime"`
	LastUpdated       string        `json:"lastUpdated"`
}

type SeriesStatus struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	RecordType string `json:"recordType"`
}

type tvdbErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
