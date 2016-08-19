package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const dbPath = "./data.db"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	debug := flag.Bool("debug", false, "enable debug output")
	flag.Parse()

	if *debug != false {
		log.SetOutput(os.Stdout)
	}
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
