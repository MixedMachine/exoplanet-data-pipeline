# Data Ingestion Service
*Mixed Machine* <br />
*mixedmachine.dev@gmail.com*

## Description
Pulls exoplanet data from NASA's API and stores it in raw form in MongoDB. Once the data is stored, the service publishes a message on a NATS queue. Uses NASA's exoplanet data [API](https://exoplanetarchive.ipac.caltech.edu/) to retrieve data. This service will also delete from mongodb so only one service interacts with the database.

### Language
Go

## Todo
- [x] Add mongodb functionality
- [x] Add NATS functionality
- [ ] Add Dockerfile
- [ ] Add Kubernetes deployment files
- [ ] Add Helm chart

**Note:** This service is not meant to be run on its own. It is meant to be run as a part of the larger system. This service is also in proof of concept stage so it's very messy and will be refactored once the system is more complete.