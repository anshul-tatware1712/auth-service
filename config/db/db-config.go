package config

import (
	"auth-service/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func CreateDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
    cfg.User = config.GetString("DBUSER")
    cfg.Net = config.GetString("NET")
    cfg.Addr = config.GetString("ADDR")
	cfg.DBName = config.GetString("DBNAME")

	var err error
    var db *sql.DB
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
    return db, nil
}