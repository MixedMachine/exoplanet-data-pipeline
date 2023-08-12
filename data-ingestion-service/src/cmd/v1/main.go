package main

import (
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/api"
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/database"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		PadLevelText: true,
	})

	// TODO: Add log level as a flag
	// log.SetLevel(log.DebugLevel)
	// TODO: Add log file as a flag
	// log.SetReportCaller(true)

	log.Info("Starting data ingestion service...")

	// TODO: break out into separate functions (e.g. api, database, messaging)
	client := api.NewExoplanetArchive()
	query := api.NewQueryBuilder().
		AddSelect("*").
		AddFrom("k2pandc").
		AddWhere().
		AddWhereParameter("rowupdate", ">", "2023-07-01").
		AddFormat("json").
		Build()
	data, err := client.GetExoplanets(query)
	if err != nil {
		panic(err)
	}

	mongoClient := database.ConnectDB("mongodb://root:root@localhost:27017")
	mongoCollection := database.GetCollection(mongoClient, "exoplanets", "k2pandc")

	inserted := []string{}
	for _, planet := range *data {
		log.Infof("Inserting: %v, %v (%v)",
			planet["pl_name"], planet["hostname"], planet["disc_year"])
		id, err := database.InsertOne(mongoCollection, planet)
		if err != nil {
			log.Warn(err)
		} else {
			log.Info("Inserted id: ", id)
			inserted = append(inserted, id)
		}
		log.Info("--------------------")
	}
	log.Infof("Inserted %v planets\n", len(inserted))
}
