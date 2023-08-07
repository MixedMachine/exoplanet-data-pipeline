package test

import (
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/api"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExoplanetArchive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Define your mock response
	mockResponse := `[{"pl_name": "Exoplanet1", "disposition": "Confirmed", "disc_year": 2005, "rowupdate": "2021-10-01"}]`
	httpmock.RegisterResponder("GET", "https://exoplanetarchive.ipac.caltech.edu/TAP/sync?query=your_query_here",
		httpmock.NewStringResponder(200, mockResponse))

	// Instantiate ExoplanetArchive
	e := api.NewExoplanetArchive()

	t.Run("Test GetExoplanets", func(t *testing.T) {
		exoplanets, err := e.GetExoplanets("?query=your_query_here")
		assert.NoError(t, err)

		assert.Equal(t, "Exoplanet1", (*exoplanets)[0].Name)
		assert.Equal(t, "Confirmed", (*exoplanets)[0].Disposition)
		assert.Equal(t, 2005, (*exoplanets)[0].DiscoveryYear)
		assert.Equal(t, "2021-10-01", (*exoplanets)[0].Updated)
	})
}
