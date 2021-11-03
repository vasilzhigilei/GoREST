package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
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
}

func createCustomer(w http.ResponseWriter, r *http.Request){
	// note to self, hash password
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	checkErrHttp(err, true, &w)
	
	err = db.CreateCustomer(&customer)
	checkErrHttp(err, true, &w)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request){

}

func getCertificates(w http.ResponseWriter, r *http.Request){
	val := chi.URLParam(r, "customer_id")
	checkErrHttp(nil, len(val) > 0, &w)
	
	certificates, err := db.GetCertificates(val)
	checkErrHttp(err, true, &w)

	jsonResp, err := json.Marshal(certificates)
	checkErrHttp(err, true, &w)

	w.Write(jsonResp)
}

func createCertificate(w http.ResponseWriter, r *http.Request){

}

func toggleCertificate(w http.ResponseWriter, r *http.Request){

}

func deleteCertificate(w http.ResponseWriter, r *http.Request){

}