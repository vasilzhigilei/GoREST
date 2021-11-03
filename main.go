package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func dbSetup() (*Database, error) {
	dbUser := os.Getenv("GOREST_POSTGRES_USER")
	dbPassword := os.Getenv("GOREST_POSTGRES_PASSWORD")
	dbURL := os.Getenv("GOREST_POSTGRES_URL") // without 5432
	dbName := os.Getenv("GOREST_POSTGRES_NAME")

	database, err := InitializeDB(dbUser, dbPassword, dbURL, dbName)
	
	return database, err
}

func main() {
	// ROUTER SETUP
	r := chi.NewRouter()
	
	// Logs start/end, response status, and more of each request
	r.Use(middleware.Logger)
	// Timeout (60 second timeout)
	r.Use(middleware.Timeout(60*time.Second))

	// Full API structure
	// base_url/customers/{customer_id}/certificates/{certificate_id}
	// Features
	//	/customers
	//		* Create a customer
	//		* Delete a customer
	//	/customers/{customer_id}
	//		* Delete a customer
	//	/customers/{customer_id}/certificates
	//		* Get all active certificates for a customer
	//		* Create new certificate
	//	/customers/{customer_id}/certificates/{certificate_id}
	//		* Update active status for a certificate
	//		* Delete certificate

	// certificate IDs are relative to the customer
	r.Route("/customers", func(r chi.Router) {
		r.Post("/", createCustomer)
		r.Route("/{customer_id}/certificates", func (r chi.Router)  {
			r.Delete("/", deleteCustomer)
			r.Route("/certificates", func (r chi.Router)  {
				r.Get("/", getCertificates)
				r.Post("/", createCertificate)
				r.Route("/{certificate_id}", func(r chi.Router) {
					r.Put("/", toggleCertificate) // updates activated status, details in sent json
					r.Delete("/", deleteCertificate)
				})
			})
		})
	})
	
	http.ListenAndServe(":3000", r)
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}