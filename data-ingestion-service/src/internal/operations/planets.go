package operations

import (
	"encoding/json"
	"time"

	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/database"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	NATS_CHANNEL_INGESTED = "exoplanets.ingested"
	EXOPLANET_NAME        = "pl_name"
	EXOPLANET_HOST        = "hostname"
	EXOPLANET_YEAR        = "disc_year"
	DIV                   = "--------------------"
)

func SavePlanets(mongoCollection *mongo.Collection, natsClient *nats.Conn, planets []map[string]any) {
	inserted := []string{}
	for _, planet := range planets {
		log.Infof("Inserting: %v, %v (%v)",
			planet[EXOPLANET_NAME], planet[EXOPLANET_HOST], planet[EXOPLANET_YEAR])

		id, err := database.InsertOne(mongoCollection, planet)
		if err != nil {
			log.Warn(err)
		} else {
			log.Info("Inserted id: ", id)
			inserted = append(inserted, id)
			natsClient.Publish(NATS_CHANNEL_INGESTED, []byte(id))
		}
		log.Info(DIV)

	}
	log.Infof("Inserted %v planets", len(inserted))
}

func CleanUpPlanets(mongoCollection *mongo.Collection, sub *nats.Subscription) {
	msg, _ := sub.NextMsg(10 * time.Millisecond)
	if msg != nil {
		messageBody := map[string]string{}
		json.Unmarshal(msg.Data, &messageBody)
		log.Infof("Received from %s | %v", msg.Subject, messageBody)
		cleanUpPlanet(messageBody["_id"], mongoCollection)
	} else {
		print("...\r")
	}
}

func cleanUpPlanet(id string, mongoCollection *mongo.Collection) {
	if id != "" { //&& msgData["status"] == COMPLETE {
		log.Info("Deleting ", id, "...")
		err := database.DeleteById(mongoCollection, id)
		if err != nil {
			log.Warn(err)
		}
	}
}
