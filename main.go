package main

import (
	"github.com/kiran/NawabiHouse/connection"
	"github.com/kiran/NawabiHouse/handler"
	"log"
	"net/http"
)

func main() {
	// Connect to DB
	//if err := connection.ConnectDB(); err != nil {
	//	log.Fatal(err)
	//}
	connection.InitDB()
	log.Println("connecting to database")
	http.HandleFunc("/orders", handler.OrderHandler)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
