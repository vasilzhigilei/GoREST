package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Certificates []Certificate `json:certificates`
}

type Certificate struct {
	ID uint `json:"id"`
	Active bool `json:"active"`
	PrivateKey string `json:"privatekey"`
	Body string `json:"body"`
	WebhookURL string `json:"webhookurl"`
}

type CertificateWebhook struct {
	ID uint `json:"id"`
	Active bool `json:"active"`
}

type Active struct {
	Active bool `json:"active"`
}

// Creates a customer in the database based on customer JSON object
func createCustomer(w http.ResponseWriter, r *http.Request){
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	checkErrHttp(err, true, &w)
	
	passBytes, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 14)
	checkErrHttp(err, true, &w)

	customer.Password = string(passBytes)

	err = db.CreateCustomer(&customer)
	checkErrHttp(err, true, &w)
}

// Deletes a customer in the database based on the included ID
func deleteCustomer(w http.ResponseWriter, r *http.Request){
	customer_id := chi.URLParam(r, "customer_id")
	checkErrHttp(nil, len(customer_id) > 0, &w)

	err := db.DeleteCustomer(customer_id)
	checkErrHttp(err, true, &w)
}

// Provides all active certificates for a given customer based on customer ID
func getCertificates(w http.ResponseWriter, r *http.Request){
	customer_id := chi.URLParam(r, "customer_id")
	checkErrHttp(nil, len(customer_id) > 0, &w)
	
	certificates, err := db.GetCertificates(customer_id)
	checkErrHttp(err, true, &w)

	var activeCerts []Certificate
	for i := 0; i < len(certificates); i++ {
		if certificates[i].Active {
			activeCerts = append(activeCerts, certificates[i])
		}
	}

	jsonResp, err := json.Marshal(activeCerts)
	checkErrHttp(err, true, &w)

	w.Write(jsonResp)
}

// Adds additional certificate to customer based on customer ID
func createCertificate(w http.ResponseWriter, r *http.Request){
	customer_id := chi.URLParam(r, "customer_id")
	checkErrHttp(nil, len(customer_id) > 0, &w)

	var certificate Certificate
	err := json.NewDecoder(r.Body).Decode(&certificate)
	checkErrHttp(err, true, &w)

	err = db.CreateCertificate(customer_id, &certificate)
	checkErrHttp(err, true, &w)
}

// Either makes certificate with ID=certificate_id active or unactive
// If certificate field for URL is not empty, sends event update to URL
func toggleCertificate(w http.ResponseWriter, r *http.Request){
	customer_id := chi.URLParam(r, "customer_id")
	checkErrHttp(nil, len(customer_id) > 0, &w)
	
	certificate_id := chi.URLParam(r, "certificate_id")
	checkErrHttp(nil, len(certificate_id) > 0, &w)
	certificate_id_uint, err := strconv.ParseUint(certificate_id, 10, 0)
	checkErrHttp(err, true, &w)

	var active Active
	err = json.NewDecoder(r.Body).Decode(&active)
	checkErrHttp(err, true, &w)
	
	err = db.ToggleCertificate(customer_id, uint(certificate_id_uint), active)
	checkErrHttp(err, true, &w)
}