# GoFlow 🌀
GoFlow is a hands-on microservices-based project developed using Go. It's designed to help understand and implement real-world patterns in distributed systems — including service communication, message queues, containerization, and database integration.

## 🌟 Overview
GoFlow is composed of multiple microservices that work together in a containerized environment. Each service handles a specific responsibility and communicates with others either directly (via HTTP) or indirectly (via a message queue). The project is Dockerized and orchestrated using Docker Compose for a smooth local development experience.

## 🧩 Services Overview
### 🔐 Authentication Service
Handles user authentication and management.

Connected to a PostgreSQL database for storing user credentials and session data.

### 🔀 Broker Service
Optional single point of entry.

Forwards incoming client requests to the appropriate microservice.

### 📧 Mail Service
Uses Mailhog for local email testing.

Sends and previews emails when triggered.

### 📝 Logger Service
Captures and stores logs from other services.

Persists log data in MongoDB.

### 🎧 Listener Service
Subscribes to events from RabbitMQ.

Processes incoming messages asynchronously, demonstrating event-driven architecture.

### 🛠️ Tech Stack
Language: Go (Golang)

Databases: PostgreSQL, MongoDB

Messaging Queue: RabbitMQ

Email Testing: Mailhog

Containerization: Docker, Docker Compose

Build: Makefile

## 🧪 Getting Started

### 1. Clone the repository

`````````
git clone https://github.com/Adit0507/GoFlow.git
cd GoFlow
`````````
### 2. Build and run all services
`````````
make up_build
`````````
- Build all Go binaries
- Stop any running Docker containers
- Start all containers with fresh builds


### 3. Start services without rebuilding
`````````
make up
`````````

### 4. Stop all services 
`````````
make down
`````````

### 5. Launch the Front-end 
`````````
make start
`````````
to stop frontend
`````````
make stop
`````````



## 🚧 Planned Enhancements
gRPC Integration
Services will eventually communicate using gRPC for better performance and strong type guarantees.
