

# GoFlow 🐹🌀
GoFlow simula la orquestación de servicios del mundo real con límites claros de servicios, comunicación asíncrona y despliegues en contenedores.

El proyecto incluye múltiples servicios—cada uno con una responsabilidad específica—conectados a través de HTTP y gRPC. Utiliza RabbitMQ para colas de mensajes, PostgreSQL y MongoDB para persistencia, y Mailhog para pruebas de correo electrónico. Todo está contenerizado con Docker **o** orquestado con Kubernetes y controlado a través de un flujo de trabajo simple con Makefile.


Una interfaz de usuario front-end ligera muestra las interacciones, mientras que un flujo de trabajo limpio con Makefile (o un solo kubectl apply -f k8s/) inicia todo, brindándote una vista práctica de sistemas escalables y orientados a eventos en acción.


## 📚 Tabla de Contenidos

- 🌟[ Resumen](#-overview)
- 🧩[ Resumen de Servicios](#-services-overview)
  - [🔐 Servicio de Autenticación](#-authentication-service)
  - 🔀 [Servicio Broker](#-broker-service)
  - 📧[ Servicio de Correo](#-mail-service)
  - 📝[ Servicio de Registro](#-logger-service)
  - 🎧[ Servicio Escucha](#-listener-service)
- 🔌[ Comunicación entre Servicios](#-interservice-communication)
- 🏛️ [ Diagrama de Arquitectura ](#-architecture-diagram)
- 🛠️ [ Stack Tecnológico](#-tech-stack)
- 🧪[ Primeros Pasos](#-getting-started)
- 📦 [ Manifiestos de Kubernetes](#-kubernetes-manifests)
- 📊 [ Panel de Kubernetes](#-kubernetes-dashboard)
- 📸[ Demo](#demo)
- 🤝[ Contribuir](#-contributing)


 ## 🌟 Resumen
GoFlow es un sistema backend basado en microservicios construido con Go. Simula la orquestación de servicios del mundo real con límites claros de servicios, comunicación asíncrona y despliegues en contenedores.

El proyecto incluye múltiples servicios — cada uno con una responsabilidad específica — conectados a través de HTTP y ahora también gRPC. Utiliza herramientas como RabbitMQ para colas de mensajes, PostgreSQL y MongoDB para persistencia, y Mailhog para pruebas de correo electrónico. Todo está contenerizado usando Docker y controlado a través de un flujo de trabajo simple con Makefile.


 ## 🧩 Resumen de Servicios
 ### 🔐 Servicio de Autenticación
 Maneja la autenticación y gestión de usuarios.
 
 Conectado a una base de datos PostgreSQL para almacenar credenciales de usuarios y datos de sesión.
 
 ### 🔀 Servicio Broker
 Punto de entrada único opcional.
 
 Reenvía las solicitudes entrantes del cliente al microservicio correspondiente.
 
 ### 📧 Servicio de Correo
 Utiliza Mailhog para pruebas locales de correo electrónico.
 
 Envía y previsualiza correos electrónicos cuando se activa.
 
 ### 📝 Servicio de Registro
 Captura y almacena registros (logs) de otros servicios.
 
 Persiste los datos de registro en MongoDB.
 
 ### 🎧 Servicio Escucha
 Se suscribe a eventos de RabbitMQ.
 
 Procesa los mensajes entrantes de forma asíncrona, demostrando una arquitectura orientada a eventos.

## 🔌 Comunicación entre Servicios
GoFlow ahora soporta ambos:

Comunicación HTTP entre el Broker y los servicios

gRPC para una comunicación más rápida y fuertemente tipificada entre servicios internos


## 🏛️ Diagrama de Arquitectura 

 ![20](https://github.com/user-attachments/assets/1a8e6bb5-c279-49ce-b790-94caf63ed6e1)

 ## 🛠️ Stack Tecnológico
 <b>Lenguaje</b>: Go (Golang)
 
 <b>Bases de datos</b>: PostgreSQL, MongoDB
 
 <b>Cola de mensajes</b>: RabbitMQ
 
 <b>Pruebas de correo</b>: Mailhog
 
 <b>Contenerización</b>: Docker, Docker Compose
 
 <b>Compilación</b>: Makefile
 
 ## 🧪 Primeros Pasos
 
 ### 1. Clonar el repositorio
 
 `````` ```
 git clone https://github.com/Adit0507/GoFlow.git
 cd GoFlow
 `````` ```
 ### 2. Compilar y ejecutar todos los servicios
 `````` ```
 make up_build
 `````` ```
 - Compilar todos los binarios de Go
 - Detener cualquier contenedor Docker en ejecución
 - Iniciar todos los contenedores con compilaciones nuevas
 
 
 ### 3. Iniciar servicios sin recompilar
 `````` ```
 make up
 `````` ```
 
 ### 4. Detener todos los servicios 
 `````` ```
 make down
 `````` ```
 
 ### 5. Iniciar el Front-end 
 `````` ```
 make start
 `````` ```
 para detener el front-end
 `````` ```
 make stop
 `````` ```


## 📦 Manifiestos de Kubernetes
Cada microservicio ahora incluye un manifiesto de Deployment/Service en k8s/.
Inicia el stack completo en cualquier clúster de Kubernetes (por ejemplo, Minikube) con:

``````
kubectl apply -f k8s/
kubectl apply -f ingress.yaml
``````
- Cada <b>YAML</b> (authentication.yaml, broker.yaml, rabbit.yaml, etc.) crea un Deployment junto con su Service correspondiente, poniendo todos los componentes en línea igual que Docker Compose.

- ingress.yaml crea my-ingress, el cual enruta el tráfico desde dos nombres de host:
     - front-end.info → Servicio front-end :80

    - broker-service.info → Servicio broker-service :8080
La anotación nginx.ingress.kubernetes.io/rewrite-target: /$1 mantiene las URLs limpias.

## 📊 Panel de Kubernetes
Después de aplicar los manifiestos, verifica la salud del clúster con:

``````
minikube dashboard
``````

## 🧩 Cargas de Trabajo de Kubernetes

<img src="/assets/kubernetes-dashboard.png">

## Demo

Mostrando el funcionamiento de los microservicios de <b>GoFlow</b> en acción.

###  <i>Servicio Broker </i>
<img src="./assets/broker-service.PNG" />

### 🔒 <i>Servicio de Autenticación </i>
<img src="./assets/auth-service.PNG" />


### <i>Servicio de Registro (ocurriendo vía gRPC 🔥) </i>
<img src="./assets/logger-service.PNG" />


### 📧 <i>Servicio de Correo </i> 
<img src="./assets/mail-service.PNG" />

<br />
<b> MailHog (Para Pruebas de Correo) </b>
<img src="./assets/mailhog.PNG" />

### <i> Registro gRPC </i>
<img src="./assets/grPC-Log.PNG" />



## 🤝 Contribuir
Este es un proyecto de pasión construido para explorar la arquitectura backend. ¡Las contribuciones, sugerencias o simplemente retroalimentación son bienvenidas!
