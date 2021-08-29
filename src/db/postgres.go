package db

import (
	"embed"
	"fmt"
	migrations3 "github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	//migrations2 "github.com/go-pg/migrations"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	options := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     address,
		Database: "postgres",
		PoolSize: 50,
	}
	con = pg.Connect(options)
	if con == nil {
		log.Fatalln("cannot connect to postgres")
	}

	return con
}

func Migrations() {

	collection := migrations3.NewCollection()
	err := collection.DiscoverSQLMigrationsFromFilesystem(http.FS(migrationFiles), "migrations")
	if err != nil {
		print(err.Error())
	}
	_, _, err = collection.Run(NewDBConn(), "up")
	if err != nil {
		print(err.Error())
	}
}
