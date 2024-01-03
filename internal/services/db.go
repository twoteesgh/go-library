package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func DB() *sql.DB {
	// Database initialization
	sqlUrl := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?multiStatements=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	sqlDb, sqlErr := sql.Open("mysql", sqlUrl)
	if sqlErr != nil {
		panic(sqlErr)
	}
	if err := sqlDb.Ping(); err != nil {
		panic(err)
	}

	// Run database migrations
	if driver, err := mysql.WithInstance(sqlDb, &mysql.Config{}); err != nil {
		panic(err)
	} else if m, err := migrate.NewWithDatabaseInstance(
		"file://scripts/migrations",
		"mysql",
		driver,
	); err != nil {
		panic(err)
	} else {
		m.Up()
	}

	return sqlDb
}
