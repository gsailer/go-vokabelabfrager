package main

import (
	"log"
	"net/http"
)

const dbPath = "/usr/local/etc/data.db"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":80", router))
}
