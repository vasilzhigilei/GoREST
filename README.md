# GoREST
HTTP-based RESTful API written in Golang. Utilizes Postgres and Docker.

## Technologies used
* Golang
    * Chi
    * pgx (Postgres driver library)
* Postgres
* Docker

## Setup and run
```
sudo docker-compose up --build
```

## API details
### JSON structure examples
#### Customer
```
{
    "name": "James",
    "password": "mypassword",
    "certificates": []
}

```
#### Certificate
```
{
    "id": 3,
    "Active": true,
    "privatekey": "helloworld",
    "body": "body-of-the-certificate"
}
```
<sup>Note: Certificate ID uniqueness is not checked by the server. When toggling/deleting a certificate, the operation is only performed on the first certificate found with the provided ID.</sup>
#### Active status
```
}
    "active": false
}
```

### Full API structure

```base_url/customers/{customer_id}/certificates/{certificate_id}```

### Features

```/customers```

* Create a customer
* Delete a customer

```/customers/{customer_id}```
* Delete a customer

```/customers/{customer_id}/certificates```
* Get all active certificates for a customer
* Create new certificate

```/customers/{customer_id}/certificates/{certificate_id}```
* Update active status for a certificate

## Project structure
* main.go - contains main function
* handlers - handlers for http requests
* db.go - functions relating to database access