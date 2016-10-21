package main

import (
	"log"
	"net/http"
)

const dbPath = "/usr/local/etc/vokabelabfrager/data.db"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
