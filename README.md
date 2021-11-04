# GoREST
HTTP-based RESTful API written in Golang. Utilizes Postgres and Docker.

### Setup and run
```
sudo docker-compose up --build
```

further considerations:
future work
protection against SQL injections

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