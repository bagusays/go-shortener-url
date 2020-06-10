package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DbConn() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/urlshortener")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
