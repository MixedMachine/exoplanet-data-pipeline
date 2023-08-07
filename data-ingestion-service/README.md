# Data Ingestion Service
*Mixed Machine* <br />
*mixedmachine.dev@gmail.com*

## Description
Pulls exoplanet data from NASA's API and stores it in raw form in MongoDB. Once the data is stored, the service publishes a message on a NATS queue. Uses NASA's exoplanet data [API](https://exoplanetarchive.ipac.caltech.edu/) to retrieve data.

### Language
Go

## Todo
- [ ] Add mongodb functionality
- [ ] Add NATS functionality
- [ ] Add Dockerfile
- [ ] Add Kubernetes deployment files
- [ ] Add Helm chart
