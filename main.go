package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

func main() {
	log.Printf("Starting")
	database, err := sql.Open("sqlite3", "./interviewSupport.db")
	
	if err != nil{
		log.Fatal("Failed to open DB")
	}

	log.Printf("Preparing Table")
	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS lcproblems (
		id INTEGER PRIMARY KEY,
		title TEXT NOT NULL,
		link TEXT NOT NULL, 
		type TEXT, 
		best_soution_time TEXT,
		difficulty TEXT NOT NULL)`)
	
	if err != nil{
		log.Fatal("Failed to prepare statment")
	}

	statement.Exec()


	addLCProblem, err := database.Prepare("INSERT INTO lcproblems (id, title, link, difficulty) VALUES (?, ?, ?, ?)")

	if err != nil{
		log.Fatal("Failed to prepare statment")
	}

	addLCProblem.Exec(56, "Merge Intervals", "https://leetcode.com/problems/merge-intervals/description/", "Medium")

	rows, err := database.Query("SELECT title, link, difficulty FROM lcproblems")
	var title string
	var link string
	var difficulty string
	for rows.Next() {
		rows.Scan(&title, &link, &difficulty)
		fmt.Println(title + " | " + difficulty + " | " + link)
	}
}
