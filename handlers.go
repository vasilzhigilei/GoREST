package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	checkErrHttp(err, true, &w)
	fmt.Fprintf(w, string(body))
	//_, err := db.conn.Exec(context.Background(), "INSERT INTO customers VALUES($1, $2, $3, $4)")
}

func deleteCustomer(w http.ResponseWriter, r *http.Request){

}

func getCertificates(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	custID, ok := ctx.Value("customer_id").(uint)
	checkErrHttp(nil, ok, &w)
	querystr := "SELECT certificates FROM customers WHERE id=" + string(custID) + ";"
	var result string
	err := db.conn.QueryRow(context.Background(), querystr).Scan(result)
	checkErrHttp(err, true, &w)
	w.Write([]byte(result))
}

func createCertificate(w http.ResponseWriter, r *http.Request){

}

func toggleCertificate(w http.ResponseWriter, r *http.Request){

}

func deleteCertificate(w http.ResponseWriter, r *http.Request){

}