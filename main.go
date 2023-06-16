package main

import (
	"os"

	"github.com/mataralhawiti/gomovies/internal" // the directory Not the file!
	"github.com/mataralhawiti/gomovies/parser"   // the directory Not the file!
)

func main() {
	args := os.Args[1:]
	myMovies := parser.ParseJSON(args[0])

	// New imple
	internal.DisplayInfo(myMovies)


	// bigquery
	// projectID := "xxxxxxxxx"
	// c := bigquery.CreateBqClient(projectID)
	// fmt.Println(c)
	//2023/06/13 14:59:34 bigquery.NewClient: bigquery: constructing client: google: could not find default credentials.
	//See https://cloud.google.com/docs/authentication/external/set-up-adc for more information

	// InsertBQ := bigquery.InsertIntoBq()
	// ReadBG := bigquery.ReadFromBq()
	// print(InsertBQ, ReadBG)
}
