package main

import (
	"log"
	"os"

	"github.com/mataralhawiti/gomovies/bigquery"
	"github.com/mataralhawiti/gomovies/internal" // the directory Not the file!
	"github.com/mataralhawiti/gomovies/parser"   // the directory Not the file!
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Only pass one arg, your movies JSON file path")
	}
	myMovies := parser.ParseJSON(args[0])

	// New imple - display stats about your movies list
	internal.DisplayInfo(myMovies)

	// bigquery
	projectID := "golang-389808"
	dataSet := "dummyds"
	table := "dumtable"
	c := bigquery.CreateBqClient(projectID)
	//2023/06/13 14:59:34 bigquery.NewClient: bigquery: constructing client: google: could not find default credentials.
	//See https://cloud.google.com/docs/authentication/external/set-up-adc for more information

	// create bq dataset if not exists
	//bigquery.CreateDataSet(dataSet, c)

	// create bq table if not exists
	//bigquery.CreateTable(dataSet, table, c)

	// insert into bq table
	//bigquery.InsertIntoBq(dataSet, table, c, *myMovies)

	// query bq table
	// ReadBG := bigquery.ReadFromBq()
}
