package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Database struct {
	conn *pgx.Conn
}

func InitializeDB(username, password, URL, dbname string) (*Database, error) {
	var dbURL string = fmt.Sprintf("postgres://%s:%s@%s:5432/%s", username, password, URL, dbname)
	conn, err := pgx.Connect(context.Background(), dbURL)

	d := new(Database)
	d.conn = conn

	return d, err
}


func (d *Database) CreateCustomer(customer *Customer) error {
	_, err := db.conn.Exec(context.Background(), 
		"INSERT INTO customers (name, password, certificates) VALUES($1, $2, $3)", customer.Name, customer.Password, customer.Certificates)
	return err
}

func (d *Database) DeleteCustomer(customer_id string) error {
	_, err := db.conn.Exec(context.Background(), 
		"DELETE FROM customers WHERE id=" + customer_id + ";")
	return err
}

func (d *Database) GetCertificates(customer_id string) ([]Certificate, error) {
	querystr := "SELECT certificates FROM customers WHERE id=" + customer_id + ";"
	var certificates []Certificate
	err := db.conn.QueryRow(context.Background(), querystr).Scan(&certificates)
	return certificates, err
}

func (d *Database) CreateCertificate(customer_id string, certificate *Certificate) error {
	jsonCert, err := json.Marshal(certificate)
	if err != nil {
		return err
	}
	_, err = db.conn.Exec(context.Background(), 
		"UPDATE customers SET certificates=certificates||'" + string(jsonCert) + "'::jsonb WHERE id=" + customer_id + ";")
	return err
}

func (d *Database) ToggleCertificate(customer_id string, certificate_id uint, active bool) error {
	querystr := "SELECT certificates FROM customers WHERE id=" + customer_id + ";"
	var certificates []Certificate
	err := db.conn.QueryRow(context.Background(), querystr).Scan(&certificates)
	if err != nil {
		return err
	}
	for i := 0; i < len(certificates); i++ {
		if certificates[i].ID == certificate_id {
			certificates[i].Active = active
			jsonCert, err := json.Marshal(certificates)
			if err != nil {
				return err
			}
			_, err = db.conn.Exec(context.Background(), 
				"UPDATE customers SET certificates='" + string(jsonCert) + "'::jsonb WHERE id=" + customer_id + ";")
			return err
		}
	}
	return err
}