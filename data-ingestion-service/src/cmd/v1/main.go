package main

import (
	"os"
	"time"

	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/internal/operations"
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/api"
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/database"
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/messaging"

	log "github.com/sirupsen/logrus"
)

const (
	MONGO_URI              = "mongodb://root:root@localhost:27017"
	MONGO_DB               = "exoplanets"
	MONGO_COLLECTION       = api.EXOPLANET_ARCHIVE_FROM
	NATS_URI               = "http://localhost:4222"
	NATS_CHANNEL_INGESTED  = "exoplanets.ingested"
	NATS_CHANNEL_PROCESSED = "exoplanets.processed"
	COMPLETE               = "complete"
	LOG_FILE_KEY		   = "LOG_FILE"
	LOG_LEVEL_KEY		  = "LOG_LEVEL"
	DEBUG				   = "DEBUG"
	INFO				   = "INFO"
	WARN				   = "WARN"
	ERROR				   = "ERROR"
)

var (
	startDate = "2023-06-01"
	throughDate   = "2023-08-30"
	sleepTime = 500 * time.Millisecond
)

func main() {
	initializeLogger()

	log.Info("Starting data ingestion service...")

	client := api.NewExoplanetArchive()
	query := api.BuildQueryBetween(startDate, throughDate)
	data, err := client.GetExoplanets(query)
	if err != nil {
		log.Error(err)
	}

	mongoManager := database.NewDatabaseManager(MONGO_URI, MONGO_DB, MONGO_COLLECTION)

	natsManager := messaging.NewNatsManager(NATS_URI)

	sub := natsManager.Subscribe(NATS_CHANNEL_PROCESSED)
	defer natsManager.Close()

	operations.SavePlanets(mongoManager.GetCollection(), natsManager.GetClient(), *data)

	for {
		operations.CleanUpPlanets(mongoManager.GetCollection(), sub)
		time.Sleep(sleepTime)
	}

}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func initializeLogger() {
	log.SetFormatter(&log.TextFormatter{
		PadLevelText: true,
	})
	logLevel := GetEnv(LOG_LEVEL_KEY, INFO)
	switch logLevel {
	case DEBUG:
		log.SetLevel(log.DebugLevel)
	case INFO:
		log.SetLevel(log.InfoLevel)
	case WARN:
		log.SetLevel(log.WarnLevel)
	case ERROR:
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}

	logFile := GetEnv(LOG_FILE_KEY, "")
	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Info("Failed to log to file, using default stderr")
		}
	}
}
