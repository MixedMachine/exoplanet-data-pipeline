# Exoplanet Data Processing Service
*Mixed Machine* <br />
*mixedmachine.dev@gmail.com*

## Description
This project provides a data processing pipeline for exoplanet data. It uses a microservice architecture with individual services built with Python, Go, and Java. Services interact through a NATS messaging system, and data is stored in MongoDB and PostgreSQL databases. The entire system is containerized using Docker and orchestrated with Kubernetes.

## Architecture
Here's a brief overview of the services in this system:

- Data Ingestion Service (Go): Pulls exoplanet data from NASA's API and stores it in raw form in MongoDB. Once the data is stored, the service publishes a message on a NATS queue.

- Data Processing Service (Python): Listens for messages from the Data Ingestion Service. When it receives a message, it pulls the corresponding raw data from MongoDB, processes it, and stores the results in PostgreSQL.

- Data Deletion Service (Go): Listens for completion messages from the Data Processing Service. When it receives a message, it deletes the raw data from MongoDB.

- Notification Service (Javascript / Node.js): Listens for completion messages from the Data Processing Service. When it receives a message, it notifies the user (e.g., by email) that their data is ready.

- API Service (Java / Spring Boot): Provides a RESTful API for users to retrieve processed data from PostgreSQL.

## Setup

The project uses Docker for local development and Kubernetes for production. Detailed setup instructions will be provided in the respective directories.

- `./data-ingestion-service`
- `./data-processing-service`
- `./data-deletion-service`
- `./notification-service`
- `./api-service`

use the make file to deploy, or use the kubernetes files to deploy to a cluster

## Contributing

As the project progresses, we welcome contributions! Check out our open issues or propose new features or improvements by creating a new issue.

---

## Contact

For questions or feedback, please open an issue on this repo and we'll respond as soon as we can.

## License:
This project is licensed under the MIT License - see the 
[LICENSE.md](./LICENSE.txt) file for details.
