package api

import (
	// "github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/internal/models"

	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	// eventual refactor: put in env var
	EXOPLANET_ARCHIVE_BASE_URL = "https://exoplanetarchive.ipac.caltech.edu/TAP/sync"
	SELECT                     = "*"
	EXOPLANET_ARCHIVE_FROM     = "k2pandc"
	ROW_UPDATE_FIELD           = "rowupdate"
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

func decodeJSON(body []byte, data *[]map[string]any) error {
	err := json.Unmarshal(body, data)
	if err != nil {
		log.Errorf("Error decoding JSON: %v", err)
		return err
	}

	return nil
}

func (e *ExoplanetArchive) GetExoplanets(query string) (*[]map[string]any, error) {
	data := new([]map[string]any)

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

func BuildQueryBetween(startDate, endDate string) string {
	return NewQueryBuilder().
		AddSelect(SELECT).
		AddFrom(EXOPLANET_ARCHIVE_FROM).
		AddWhere().
		AddWhereParameter(ROW_UPDATE_FIELD, ">", startDate).
		AddAndWhereParameter(ROW_UPDATE_FIELD, "<=", endDate).
		AddFormat("json").
		Build()
}
