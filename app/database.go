package app

import (
	"database/sql"
	"os"
	"simplify-go/helper"

	"time"
)

func NewDB() *sql.DB {
	root := os.Getenv("SQLROOT")
	password := os.Getenv("SQLPASSWORD")
	host := os.Getenv("SQLHOST")
	port := os.Getenv("SQLPORT")
	database := os.Getenv("SQLDATABASE")

	db, err := sql.Open("mysql", root+":"+password+"@tcp("+host+":"+port+")/"+database)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
