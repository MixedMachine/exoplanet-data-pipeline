package test

import (
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/internal/models"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExoplanetString(t *testing.T) {
	// Create an Exoplanet object with known values
	exoplanet := &models.Exoplanet{
		Name:          "Exoplanet1",
		Disposition:   "Confirmed",
		DiscoveryYear: 2005,
		Updated:       "2021-10-01",
	}

	// Test String method
	t.Run("Test Exoplanet String Representation", func(t *testing.T) {
		expected := "Name: Exoplanet1, Disposition: Confirmed, Discovery Year: 2005, Updated: 2021-10-01"
		actual := exoplanet.String()

		assert.Equal(t, expected, actual)
	})

	// Test String method sad path
	t.Run("Test Exoplanet String Representation Sad Path", func(t *testing.T) {
		expected := "Name: Exoplanet1, Disposition: Confirmed, Discovery Year: 2005, Updated: 2021-10-02"
		actual := exoplanet.String()

		assert.NotEqual(t, expected, actual)
	})
}
