package main

import (
    "encoding/csv"
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func createTables() {
    db, err := sql.Open("sqlite3", dbPath)
	checkErr(err)
	defer db.Close()
	db.Exec('CREATE TABLE IF NOT EXISTS data(VID INTEGER PRIMARY KEY AUTOINCREMENT, bedeutung1 TEXT,bedeutung2 TEXT, src INTEGER)')
	db.Exec('CREATE TABLE IF NOT EXISTS stats(level INTEGER, anzahl_abfragen INTEGER, vokabel INTEGER)')
	db.Exec('CREATE TABLE IF NOT EXISTS files(FID INTEGER PRIMARY KEY AUTOINCREMENT, filename TEXT, tmstp TEXT)')
}
