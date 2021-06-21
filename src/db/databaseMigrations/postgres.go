package databasemigrations

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	db, errOpen := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")

	if errOpen != nil {
		log.Fatal(errOpen)
	}

	driver, errConfig := postgres.WithInstance(db, &postgres.Config{})

	if errConfig != nil {
		log.Fatal(errConfig)
	}

	m, errInstance := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)

	if errInstance != nil {
		log.Fatal(errInstance)
	}

	m.Steps(2)
}
