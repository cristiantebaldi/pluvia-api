package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DatabaseConfig struct {
	*sqlx.DB
}

func Initialize(context context.Context) DatabaseConfig {
	databaseURL := viper.GetString("database.url")

	db, err := sqlx.ConnectContext(context, "postgres", databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.DB.SetMaxOpenConns(100)  // The default is 0 (unlimited)
	db.DB.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	db.DB.SetConnMaxLifetime(0) // 0, connections are reused forever.

	runMigrations()

	return DatabaseConfig{db}
}

// RunMigrations run scripts on path database/migrations
func runMigrations() {
	databaseURL := viper.GetString("database.url")
	m, err := migrate.New("file://database/migrations", databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}