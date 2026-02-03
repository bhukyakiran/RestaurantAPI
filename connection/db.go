package connection

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {

	ConnStr := "host=localhost port=5432 user=kiranbhukya password=Kiran@12345 dbname=nawabihouse sslmode=disable"

	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	DB = db
	fmt.Println("connected to the database ")
	return nil

}
