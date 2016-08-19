package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
)

func loadLists() []string {
	var list []string
	db, err := sql.Open("sqlite3", dbPath)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("select filename from files")
	checkErr(err)
	for rows.Next() {
		var filename string
		rows.Scan(&filename)
		list = append(list, filename)
	}
	defer rows.Close()
	return list
}

func loadVokabel(set int) Sets {
	db, err := sql.Open("sqlite3", dbPath)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("select bedeutung1, bedeutung2 from data where src=?", set)
	checkErr(err)

	vokabeln := make(map[string]string)
	var answers []string
	var s Sets
	for rows.Next() {
		var (
			bed1 string
			bed2 string
		)
		err := rows.Scan(&bed1, &bed2)
		checkErr(err)
		vokabeln[bed1] = bed2
		answers = append(answers, bed2)
	}
	defer rows.Close()

	for bed1, bed2 := range vokabeln {
		var choices []string
		for i := 0; i < 3; i++ {
			choices = append(choices, answers[rand.Intn(len(answers))])
		}
		choices = append(choices, bed2)
		var set = Vokabelset{Question: bed1, Answer: bed2, Choices: choices}
		s = append(s, set)
	}
	return s
}
