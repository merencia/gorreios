package database

import (
	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
)

func Connect(databaseUrl string) *sqlx.DB {
	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil {
		log.WithError(err).Fatal("failed to open connection to database")
	}
	err = db.Ping()
	if err := db.Ping(); err != nil {
		log.WithError(err).Fatal("failed to ping database")
	}
	return db
}
