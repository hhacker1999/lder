package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.themoviedb.org/3"

type Tmdb struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string) Tmdb {
	return Tmdb{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (t Tmdb) ValidateKey() bool {
	req, err := http.NewRequest(http.MethodGet, baseURL+"/authentication", nil)
	if err != nil {
		return false
	}
	t.setAPIKey(req)

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	var auth AuthenticationResponse
	if err := json.NewDecoder(resp.Body).Decode(&auth); err != nil {
		return false
	}

	return auth.Success
}

func (t Tmdb) SearchMovies(query string, page int) (MovieSearchResponse, error) {
	if query == "" {
		return MovieSearchResponse{}, fmt.Errorf("search query cannot be empty")
	}
	if page < 1 {
		page = 1
	}

	u, err := url.Parse(baseURL + "/search/movie")
	if err != nil {
		return MovieSearchResponse{}, err
	}
	q := u.Query()
	q.Set("query", query)
	q.Set("page", fmt.Sprintf("%d", page))
	q.Set("include_adult", "false")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return MovieSearchResponse{}, err
	}
	t.setAPIKey(req)

	var result MovieSearchResponse
	if err := t.doJSON(req, &result); err != nil {
		return MovieSearchResponse{}, err
	}
	return result, nil
}

func (t Tmdb) MovieDetails(movieID int) (MovieDetails, error) {
	if movieID <= 0 {
		return MovieDetails{}, fmt.Errorf("invalid movie id")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/movie/%d", baseURL, movieID), nil)
	if err != nil {
		return MovieDetails{}, err
	}
	t.setAPIKey(req)
	t.setAppendToResponse(req, "credits")

	var details MovieDetails
	if err := t.doJSON(req, &details); err != nil {
		return MovieDetails{}, err
	}
	return details, nil
}

func (t Tmdb) SearchTV(query string, page int) (TVSearchResponse, error) {
	if query == "" {
		return TVSearchResponse{}, fmt.Errorf("search query cannot be empty")
	}
	if page < 1 {
		page = 1
	}

	u, err := url.Parse(baseURL + "/search/tv")
	if err != nil {
		return TVSearchResponse{}, err
	}
	q := u.Query()
	q.Set("query", query)
	q.Set("page", fmt.Sprintf("%d", page))
	q.Set("include_adult", "false")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return TVSearchResponse{}, err
	}
	t.setAPIKey(req)

	var result TVSearchResponse
	if err := t.doJSON(req, &result); err != nil {
		return TVSearchResponse{}, err
	}
	return result, nil
}

func (t Tmdb) TVExternalIDs(tvID int) (TVExternalIDs, error) {
	if tvID <= 0 {
		return TVExternalIDs{}, fmt.Errorf("invalid tv id")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/tv/%d/external_ids", baseURL, tvID), nil)
	if err != nil {
		return TVExternalIDs{}, err
	}
	t.setAPIKey(req)

	var ids TVExternalIDs
	if err := t.doJSON(req, &ids); err != nil {
		return TVExternalIDs{}, err
	}
	return ids, nil
}

func (t Tmdb) TVDetails(tvID int) (TVDetails, error) {
	if tvID <= 0 {
		return TVDetails{}, fmt.Errorf("invalid tv id")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/tv/%d", baseURL, tvID), nil)
	if err != nil {
		return TVDetails{}, err
	}
	t.setAPIKey(req)
	t.setAppendToResponse(req, "credits")

	var details TVDetails
	if err := t.doJSON(req, &details); err != nil {
		return TVDetails{}, err
	}
	return details, nil
}

func (t Tmdb) setAPIKey(req *http.Request) {
	q := req.URL.Query()
	q.Set("api_key", t.apiKey)
	req.URL.RawQuery = q.Encode()
}

func (t Tmdb) setAppendToResponse(req *http.Request, value string) {
	q := req.URL.Query()
	q.Set("append_to_response", value)
	req.URL.RawQuery = q.Encode()
}

func (t Tmdb) doJSON(req *http.Request, dest interface{}) error {
	resp, err := t.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr tmdbErrorResponse
		if json.Unmarshal(body, &apiErr) == nil && apiErr.StatusMessage != "" {
			return fmt.Errorf("tmdb api error (%d): %s", resp.StatusCode, apiErr.StatusMessage)
		}
		return fmt.Errorf("tmdb api error: status %d", resp.StatusCode)
	}

	if err := json.Unmarshal(body, dest); err != nil {
		return fmt.Errorf("decode tmdb response: %w", err)
	}
	return nil
}
