package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type connection struct {
	DB       *sql.DB
	sqlError error
}

var con connection

func Load() {

	envs := LoadDatabaseCreds()

	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s", envs["DB_USER"], envs["DB_PASS"], envs["DB_ADDRESS"], envs["DB_PORT"], envs["DB_DATABASE"])

	con.DB, con.sqlError = sql.Open(envs["DB_NAME"], source)

	if con.sqlError != nil {
		log.Fatalln("Error connecting to mySql database", con.sqlError)
	}
	// DB Options
	con.DB.SetConnMaxLifetime(3 * time.Minute)
	con.DB.SetMaxOpenConns(10)
	con.DB.SetMaxIdleConns(10)
}
