package main

import "github.com/jackc/pgx/v4"

type Database struct {
	conn *pgx.Conn
}