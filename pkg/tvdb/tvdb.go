package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const baseURL = "https://api4.thetvdb.com/v4"

type Tvdb struct {
	apiKey     string
	httpClient *http.Client

	mu    sync.Mutex
	token string
}

func New(apiKey string) Tvdb {
	return Tvdb{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (t *Tvdb) ValidateKey() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.token != "" {
		return true
	}
	token, err := t.login()
	if err != nil {
		return false
	}
	t.token = token
	return true
}

func (t *Tvdb) SeriesDetails(seriesID int) (SeriesDetails, error) {
	if seriesID <= 0 {
		return SeriesDetails{}, fmt.Errorf("invalid series id")
	}

	token, err := t.ensureToken()
	if err != nil {
		return SeriesDetails{}, err
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/series/%d", baseURL, seriesID), nil)
	if err != nil {
		return SeriesDetails{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	var resp SeriesDetailsResponse
	if err := t.doJSON(req, &resp); err != nil {
		return SeriesDetails{}, err
	}
	if resp.Status != "success" {
		return SeriesDetails{}, fmt.Errorf("tvdb api returned status %q", resp.Status)
	}
	return resp.Data, nil
}

func (t *Tvdb) ensureToken() (string, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.token != "" {
		return t.token, nil
	}
	token, err := t.login()
	if err != nil {
		return "", err
	}
	t.token = token
	return t.token, nil
}

func (t *Tvdb) login() (string, error) {
	payload, err := json.Marshal(loginRequest{APIKey: t.apiKey})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+"/login", bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	var resp loginResponse
	if err := t.doJSON(req, &resp); err != nil {
		return "", err
	}
	if resp.Status != "success" || resp.Data.Token == "" {
		return "", fmt.Errorf("tvdb login failed")
	}
	return resp.Data.Token, nil
}

func (t *Tvdb) doJSON(req *http.Request, dest interface{}) error {
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
		var apiErr tvdbErrorResponse
		if json.Unmarshal(body, &apiErr) == nil && apiErr.Message != "" {
			return fmt.Errorf("tvdb api error (%d): %s", resp.StatusCode, apiErr.Message)
		}
		return fmt.Errorf("tvdb api error: status %d", resp.StatusCode)
	}

	if err := json.Unmarshal(body, dest); err != nil {
		return fmt.Errorf("decode tvdb response: %w", err)
	}
	return nil
}
