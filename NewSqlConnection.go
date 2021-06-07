package helpers

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var sqlErr error
var beginOfMonth string

func NewSqlConnection() {

	envs := utils.LoadEnv()

	Db, sqlErr = sql.Open(
		envs["DB_NAME"],
		envs["DB_USER"]+":"+envs["DB_PASS"]+
			"@tcp("+envs["DB_ADDRESS"]+":"+envs["DB_PORT"]+")/"+envs["DB_DATABASE"]) // FROM ENV (probs should use sprintf)

	if sqlErr != nil {
		log.Fatalln("Error connecting to mySql database", sqlErr)
	}
	// DB Options
	Db.SetConnMaxLifetime(3 * time.Minute)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	BeginOfMonth()
}

func BeginOfMonth() {
	d := Db.QueryRow("SELECT CONCAT(DATE_ADD(CURDATE(),INTERVAL - DAY(CURDATE())+1 DAY)) AS date")

	err := d.Scan(&beginOfMonth)

	if err != nil {
		log.Fatalln(err)
	}

}
