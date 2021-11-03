package main

import "net/http"

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
	
}

func deleteCustomer(w http.ResponseWriter, r *http.Request){

}

func getCertificates(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	_, ok := ctx.Value("customer_id").(uint)
	if !ok {
		http.Error(w, http.StatusText(500), 500)
	}
	
}

func createCertificate(w http.ResponseWriter, r *http.Request){

}

func toggleCertificate(w http.ResponseWriter, r *http.Request){

}

func deleteCertificate(w http.ResponseWriter, r *http.Request){

}