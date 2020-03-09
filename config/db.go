package config

import (
	"database/sql"
	"log"

	//driver sql
	_ "github.com/go-sql-driver/mysql"
)

// GetConnection conectar con la base de datos
func GetConnection() *sql.DB {
	dsn := "root:rootAdmin2020@tcp(127.0.0.1:3306)/datos?charset=utf8&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
