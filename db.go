package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Database struct {
	conn *pgx.Conn
}

func InitializeDB(username, password, URL, dbname string) (*Database, error) {
	var dbURL string = fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, password, URL, dbname)
	conn, err := pgx.Connect(context.Background(), dbURL)
	
	defer conn.Close(context.Background())

	d := new(Database)
	d.conn = conn

	return d, err
}


func (d *Database) CreateCustomer(customer *Customer) error {
	_, err := d.conn.Exec(context.Background(), "execstringhere")
	return err
}

func (d *Database) DeleteCustomer(customer_id uint) error {
	return nil
}

func (d *Database) GetCertificates(customer_id uint) error {
	return nil
}

func (d *Database) CreateCertificate(customer_id uint, certificate *Certificate) error {
	return nil
}

func (d *Database) DeleteCertificate(customer_id uint, certificate_id uint) error {
	return nil
}

func (d *Database) ToggleCertificate(customer *Customer) error {
	return nil
}