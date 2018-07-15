package mysql

import (
	"database/sql"
	"log"

	// Initializes MySQL resouces.
	_ "github.com/go-sql-driver/mysql"
	"github.com/noteshare/config"
)

// database contains the active MySQL connection pool.
var database *sql.DB

// init initializes the database connection.
func init() {

	var err error
	database, err = sql.Open("mysql", config.MySQLToFormatDNS())
	if err != nil {
		log.Fatal("==> Error in library/mysql: " + err.Error())
	}

	database.SetMaxOpenConns(20)
	database.SetMaxIdleConns(20)

}

// Free releases all the database connection resources.
func Free() {
	database.Close()
}

// Open gets the database handler to the mysql server.
func Open() *sql.DB {
	return database
}
