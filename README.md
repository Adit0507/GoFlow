# GoFlow 🐹🌀
GoFlow simulates real-world service orchestration with clean service boundaries, asynchronous communication, and containerized deployments.

The project includes multiple services—each with a focused responsibility—connected through HTTP and gRPC. It leverages RabbitMQ for message queues, PostgreSQL and MongoDB for persistence, and Mailhog for email testing. Everything is containerized with Docker **or** orchestrated with Kubernetes and controlled through a simple Makefile workflow.


A lightweight front-end UI showcases the interactions, while a clean Makefile workflow (or a single kubectl apply -f k8s/) spins everything up, giving you a practical view of scalable, event-driven systems in action.


## 📚 Table of Contents

- 🌟[ Overview](#-overview)
- 🧩[Services Overview](#-services-overview)
  - [🔐 Authentication Service](#-authentication-service)
  - 🔀 [Broker Service](#-broker-service)
  - 📧[ Mail Service](#-mail-service)
  - 📝[ Logger Service](#-logger-service)
  - 🎧[ Listener Service](#-listener-service)
- 🔌[ Interservice Communication](#-interservice-communication)
- 🏛️ [ Architecture Diagram ](#-architecture-diagram)
- 🛠️ [Tech Stack](#-tech-stack)
- 🧪[ Getting Started](#-getting-started)
- 📦 [Kubernetes Manifests](#-kubernetes-manifests)
- 📊 [Kubernetes Dashboard](#-kubernetes-dashboard)
- 📸[ Demo](#demo)
- 🤝[ Contributing](#-contributing)


 ## 🌟 Overview
GoFlow is a microservices-based backend system built using Go. It simulates real-world service orchestration with clean service boundaries, asynchronous communication, and containerized deployments.

The project includes multiple services — each with a focused responsibility — connected through HTTP and now also gRPC. It leverages tools like RabbitMQ for message queues, PostgreSQL and MongoDB for persistence, and Mailhog for email testing. Everything is containerized using Docker and controlled through a simple Makefile workflow.


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

## 🔌 Interservice Communication
GoFlow now supports both:

HTTP communication between Broker and services

gRPC for faster, strongly-typed communication between internal services


## 🏛️ Architecture Diagram 

 ![20](https://github.com/user-attachments/assets/1a8e6bb5-c279-49ce-b790-94caf63ed6e1)

 ## 🛠️ Tech Stack
 <b>Language</b>: Go (Golang)
 
 <b>Databases</b>: PostgreSQL, MongoDB
 
 <b>Messaging Queue</b>: RabbitMQ
 
 <b>Email Testing</b>: Mailhog
 
 <b>Containerization</b>: Docker, Docker Compose
 
 <b>Build</b>: Makefile
 
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


## 📦 Kubernetes Manifests
Every microservice now ships with a Deployment/Service manifest in k8s/, a.
Spin up the full stack on any Kubernetes cluster (e.g., Minikube) with:

``````
kubectl apply -f k8s/
kubectl apply -f ingress.yaml
``````
- Each <b>YAML</b> (authentication.yaml, broker.yaml, rabbit.yaml, etc.) creates a Deployment plus its corresponding Service, bringing all components online just like Docker Compose.

- ingress.yaml creates my-ingress, which routes traffic from two hostnames:
     - front-end.info → front-end Service :80

    - broker-service.info → broker-service Service :8080
The nginx.ingress.kubernetes.io/rewrite-target: /$1 annotation keeps URLs clean.

## 📊 Kubernetes Dashboard
After applying the manifests, verify cluster health with:

``````
minikube dashboard
``````

## 🧩 Kubernetes Workloads

<img src="/assets/kubernetes-dashboard.png">

<i>Ik the Auth service is failing for now,fixing it 😂</i>

## Demo

Showing the working of <b>GoFlow's</b> microservices in action.

###  <i>Broker Service </i>
<img src="./assets/broker-service.PNG" />

### 🔒 <i>Authentication Service (Requires fixing 🔨) </i>
<img src="./assets/auth-service.PNG" />


### <i>Logger Service (happening via gRPC 🔥) </i>
<img src="./assets/logger-service.PNG" />


### 📧 <i>Mail Service </i> 
<img src="./assets/mail-service.PNG" />

<br />
<b> MailHog (For Email Testing) </b>
<img src="./assets/mailhog.PNG" />

### <i> gRPC Log </i>
<img src="./assets/grPC-Log.PNG" />



## 🤝 Contributing
This is a passion project built to explore backend architecture. Contributions, suggestions, or just feedback are all welcome!
