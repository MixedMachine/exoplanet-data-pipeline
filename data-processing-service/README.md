# Data Processing Service
*Mixed Machine* <br />
*mixedmachine.dev@gmail.com*

## Description
This service is responsible for processing data from the data ingest service and storing it in a sql database. It then posts that it is complete for other services.The goal is to keep only relevent data in the sql database and possible do some data processing before storing it. The additional processing can determine if the exoplanets are habitable, earth-like, or other interesting facts based on the raw data. The data in the sql service will be used by the api service to serve the data to the front end.

### Language
Python


## Todo
- [x] Add mongodb functionality
- [x] Add NATS functionality
- [x] Add SQL functionality
- [ ] Add Dockerfile
- [ ] Add Kubernetes deployment files
- [ ] Add Helm chart

**Note:** This service is not meant to be run on its own. It is meant to be run as a part of the larger system. This service is also in proof of concept stage so it's very messy and will be refactored once the system is more complete.