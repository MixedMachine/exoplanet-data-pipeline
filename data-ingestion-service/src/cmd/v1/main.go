package main

import (
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/api"
)

func main() {
	client := api.NewExoplanetArchive()
	query := api.NewQueryBuilder().
		AddSelect("pl_name,disposition,disc_year,rowupdate").
		AddFrom("k2pandc").
		AddWhere().
		AddWhereParameter("rowupdate", ">", "2023-07-01").
		AddFormat("json").
		Build()
	data, err := client.GetExoplanets(query)
	if err != nil {
		panic(err)
	}
	for _, planet := range *data {
		println(planet.String())
	}
}
