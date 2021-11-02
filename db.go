package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type Database struct {
	conn *pgx.Conn
}

func InitializeDB(username, password, baseURL, dbname string) *Database {
	var dbURL string = fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, baseURL, dbname)
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Database connection failed: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	d := new(Database)
	d.conn = conn

	return d
}


func (d *Database) CreateCustomer(customer *Customer) error {
	_, err := d.conn.Exec(context.Background(), "execstringhere")
	return err
}

func (d *Database) DeleteCustomer(customer_id uint) error {
	
}

func (d *Database) GetCertificates(customer_id uint) error {
	
}

func (d *Database) CreateCertificate(customer_id uint, certificate *Certificate) error {
	
}

func (d *Database) DeleteCertificate(customer_id uint, certificate_id uint) error {
	
}

func (d *Database) ToggleCertificate(customer *Customer) error {
	
}