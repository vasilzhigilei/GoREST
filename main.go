package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Sets up database by fetching environment variables and calling DB related function
// returns database, err
func dbSetup() (*Database, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbURL := "postgres"
	dbName := os.Getenv("POSTGRES_DB")

	database, err := InitializeDB(dbUser, dbPassword, dbURL, dbName)
	
	return database, err
}

var db *Database

func main() {
	// ROUTER SETUP
	r := chi.NewRouter()
	
	// Logs start/end, response status, and more of each request
	r.Use(middleware.Logger)
	// Timeout (60 second timeout)
	r.Use(middleware.Timeout(60*time.Second))

	// Database connection
	var err error
	db, err = dbSetup()
	if err != nil {
		for i := 0; i < 3; i++ { // try connecting 3 more times, 3 seconds apart
			time.Sleep(3*time.Second)
			db, err = dbSetup()
			if err == nil{
				break
			}
		}
		checkErr(err) // if nil nothing happens, if not, panics
	}
	defer db.conn.Close(context.Background())

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
	r.Route("/customers", func(r chi.Router) {
		r.Post("/", createCustomer)
		r.Route("/{customer_id}", func (r chi.Router)  {
			r.Delete("/", deleteCustomer)
			r.Route("/certificates", func (r chi.Router)  {
				r.Get("/", getCertificates)
				r.Post("/", createCertificate)
				r.Route("/{certificate_id}", func(r chi.Router) {
					r.Put("/", toggleCertificate)
				})
			})
		})
	})
	
	http.ListenAndServe(":8080", r)
}

// Writes HTTP error code if there is an error or if ok boolean is false
func checkErrHttp(err error, ok bool, w *http.ResponseWriter){
	if err != nil || !ok {
		http.Error(*w, http.StatusText(500), 500)
		fmt.Println(err, ok)
	}
}

// Exit the program if there is an error (used rarely)
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}