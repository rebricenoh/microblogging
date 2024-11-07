Microblogging App
================

Una aplicación de microblogging simplificada en Golang, donde los usuarios pueden publicar tweets, seguir a otros usuarios y ver un timeline de los tweets de los usuarios a los que siguen.


## Requisitos

- Docker
- Docker Compose


## Configuración del Entorno

* Agregar las siguientes variables de entorno en un archivo .env en la raíz del proyecto:

    ```bash
    DB_USER: postgres
    DB_PASSWORD: postgres
    DB_NAME: microblogging
    DB_HOST: localhost
    DB_PORT: "5432"
    ```
* Clonar este repositorio:

    ```bash
    https://github.com/rebricenoh/microblogging.git
    cd microblogging
    ```
## Estructura del proyecto

La estructura del proyecto sigue una arquitectura hexagonal, manteniendo las responsabilidades de dominio, repositorio y servicios desacopladas:

```plaintext
microblogging/
├── cmd/                    # Punto de entrada principal
│   └── main.go
├── internal/
│   ├── domain/             # Definiciones de entidades de dominio
│   ├── handler/            # Handlers de API HTTP
│   ├── repository/         # Interacciones de persistencia
│   └── service/            # Lógica de negocio central
├── pkg/                    # Paquete compartido (configuración)
│   └── config/
├── Dockerfile
├── docker-compose.yml
└── .env                    # Variables de entorno

```


## Instrucciones para ejecutar la aplicación

Construir y levantar los contenedores con Docker Compose:

```bash
docker-compose up --build
```
Esto descargará las imágenes necesarias, construirá la imagen de la aplicación y levantará los servicios de la aplicación y PostgreSQL.

La aplicación estará disponible en http://localhost:8080

## Endpoints de la API

#### Publicar un tweet

* Endpoint: `POST /tweet`
* Body de la petición:
    ```json
    {
    "user_id": 1,
    "content": "Este es un tweet de ejemplo"
    }
    ```

#### Seguir a un usuario

* Endpoint: `POST /follow`
* Body de la petición:
    ```json
    {
    "follower_id": 1,
    "followed_id": 2
    }
    ```

#### Ver el timeline

* Endpoint: `GET /timeline?user_id={user_id}`
* Descripción: Devuelve los tweets de los usuarios seguidos por el `user_id` especificado.
## Base de Datos

La aplicación utiliza PostgreSQL como base de datos. Las tablas `tweets` y `follows` se crearán automáticamente al iniciar el contenedor gracias a la función `AutoMigrate` de GORM.
## Probar la API
Se puede probar la API utilizando herramientas como `curl` o `Postman`.

Ejemplo con `curl` para publicar un tweet:
```bash
curl -X POST http://localhost:8080/tweet -d '{"user_id":1, "content":"Mi primer tweet"}' -H "Content-Type: application/json"
```
## Testing
Para ejecutar pruebas en la aplicación, usa el siguiente comando en la raíz del proyecto:
```bash
go test ./internal/service
```
Las pruebas de los casos de uso principales están ubicadas en la carpeta `internal/service`.
## Consideraciones

* La aplicación está diseñada para ser escalable y optimizada para lecturas.
* Utilicé Golang y PostgreSQL para un rendimiento y una escalabilidad óptimos.
* La configuración de infraestructura y protocolos se especifica en el archivo `docker-compose.yml`.
## Apagar los contenedores

Para detener y eliminar los contenedores, ejecutar el siguiente comando:

```bash
docker-compose down
```