package main

import (
	// "time"

	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/api"
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/database"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

const (
	NATS_CHANNEL_INGESTED = "exoplanets.ingested"
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
		AddWhereParameter("rowupdate", ">=", "2023-04-01").
		AddAndWhereParameter("rowupdate", "<", "2023-05-01").
		AddFormat("json").
		Build()
	data, err := client.GetExoplanets(query)
	if err != nil {
		panic(err)
	}

	mongoClient := database.ConnectDB("mongodb://root:root@localhost:27017")
	mongoCollection := database.GetCollection(mongoClient, "exoplanets", "k2pandc")

	natsUri := "http://localhost:4222"

	nc, err := nats.Connect(natsUri)
	if err != nil {
		log.Fatal(err)
	}
	// sub, _ := nc.SubscribeSync("exoplanets.ingest")

	defer nc.Close()
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
			nc.Publish(NATS_CHANNEL_INGESTED, []byte(id))
		}
		log.Info("--------------------")

	}
	log.Infof("Inserted %v planets\n", len(inserted))

	// for {
	// 	msg, _ := sub.NextMsg(10 * time.Millisecond)
	// 	if msg != nil {
	// 		log.Info("Received: ", string(msg.Data))
	// 	} else {
	// 		print(".")
	// 	}
	// 	time.Sleep(100 * time.Millisecond)
	// }

}
