package main

import (
	"github.com/ArminGodiz/ce_database/twitter-simulator/api"
	"github.com/ArminGodiz/ce_database/twitter-simulator/db"
	"log"
	"net/http"
)

func main() {
	db := db.GetNewDatabase()
	api.SetTwitter(db)
	log.Fatal(http.ListenAndServe(":8080", api.GetRouter()))
}
