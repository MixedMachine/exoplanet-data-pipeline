package api

import (
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/internal/models"

	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	// eventual refactor: put in env var
	EXOPLANET_ARCHIVE_BASE_URL = "https://exoplanetarchive.ipac.caltech.edu/TAP/sync"
)

type ExoplanetArchive struct {
	baseUrl   string
	apiClient *http.Client
}

func NewExoplanetArchive() *ExoplanetArchive {
	return &ExoplanetArchive{
		baseUrl:   EXOPLANET_ARCHIVE_BASE_URL,
		apiClient: http.DefaultClient,
	}
}

func extractBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
		return nil, err
	}

	return body, nil
}

func decodeJSON(body []byte, data *[]models.Exoplanet) error {
	err := json.Unmarshal(body, data)
	if err != nil {
		log.Errorf("Error decoding JSON: %v", err)
		return err
	}

	return nil
}

func (e *ExoplanetArchive) GetExoplanets(query string) (*[]models.Exoplanet, error) {
	data := new([]models.Exoplanet)

	resp, err := e.apiClient.Get(e.baseUrl + query)
	if err != nil {
		log.Errorf("Error getting exoplanets: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := extractBody(resp)
	if err != nil {
		log.Errorf("Error extracting body: %v", err)
		return nil, err
	}

	err = decodeJSON(body, data)
	if err != nil {
		log.Errorf("Error decoding JSON: %v", err)
		return nil, err
	}

	return data, nil
}
