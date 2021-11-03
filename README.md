# GoREST
HTTP-based RESTful API written in Golang. Utilizes Postgres and Docker.

### Setup and run
```
# Start Postgres container
docker run -e GOREST_POSTGRES_USER=user -e GOREST_POSTGRES_PASSWORD=secret -e GOREST_POSTGRES_URL=localhost -e GOREST_POSTGRES_NAME=gorest
```

### Technologies used
* Golang
    * Chi
    * pgx (Postgres driver library)
* Postgres
* Docker

### Project structure
* main.go - contains main function
* handlers - handlers for http requests
* db.go - functions relating to database access