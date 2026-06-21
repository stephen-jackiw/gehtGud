/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"database/sql"

	"github.com/spf13/cobra"
	_ "github.com/mattn/go-sqlite3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init is used to initialize the SQLite DB used to save your information",
	Long: `Run init before running any other commands or operations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Opening DB Connection")
		run()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run(){
	database, err := sql.Open("sqlite3", "./interviewSupport.db")
	
	if err != nil{
		log.Fatal("Failed to open DB")
	}

	//Creating Tables (move this to init command)/////////////////////////////////
	log.Printf("Creating Tables")
	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS lcproblems (
		id INTEGER PRIMARY KEY,
		lcid INTEGER,
		title TEXT NOT NULL,
		link TEXT NOT NULL, 
		typeArray INTEGER,
		best_soution_time TEXT,
		difficulty INTEGER NOT NULL,
		description TEXT)`)
	
	if err != nil{
		log.Fatal("Failed to prepare problems statment")
	}

	statement.Exec()

	// Confidence score is 1-5.  
	// 5 is passed without running the editor (blind submit),
	// 4 is minor syntax error but no changes to logic, still passed
	// 3 is minor logical errors or heavily relied on running the code for test cases, did not pass
	// 2 major logical errors or had the right idea but could not succesfully implement in time frame
	// 1 Had little to no idea how to tackle the problem or had no direction in implementation
	//
	statement, err = database.Prepare(`CREATE TABLE IF NOT EXISTS lcattempts (
		id INTEGER PRIMARY KEY,
		problemid INTEGER NOT NULL,
		attemptedDate DATETIME NOT NULL,
		timeToCompleteInMinutes INTEGER,
		pathToWriteUp TEXT, 
		timeComplexity TEXT,
		spaceComplexity TEXT,
		passed BOOLEAN NOT NULL,
		confidenceScore INTEGER,
		FOREIGN KEY(problemid) REFERENCES lcproblems(id)
	)`)
	
	if err != nil{
		fmt.Printf("FAILURE: %v\n", err)
		log.Fatal("Failed to prepare attempt statment")
	}

	statement.Exec()

	rows, err := database.Query(`
	SELECT name, sql
	FROM sqlite_schema
	WHERE type='table' AND name NOT LIKE 'sqlite_%'
`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name, sqlText string

		err := rows.Scan(&name, &sqlText)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Table: %s\nSchema: %s\n\n", name, sqlText)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
