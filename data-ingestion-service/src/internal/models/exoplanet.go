package models

import (
	"fmt"
)

type Exoplanet struct {
	Name          string `json:"pl_name"`
	Disposition   string `json:"disposition"`
	DiscoveryYear int    `json:"disc_year"`
	Updated       string `json:"rowupdate"`
}

func (e *Exoplanet) String() string {
	return fmt.Sprintf("Name: %s, Disposition: %s, Discovery Year: %d, Updated: %s",
		e.Name, e.Disposition, e.DiscoveryYear, e.Updated)
}
