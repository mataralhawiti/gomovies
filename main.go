package main

import (
	"flag"
	"log"
	"os"

	"github.com/mataralhawiti/gomovies/bigquery" // the directory Not the file!
	"github.com/mataralhawiti/gomovies/internal" // the directory Not the file!
	"github.com/mataralhawiti/gomovies/parser"   // the directory Not the file!
)

const filePath = "async_movies_full.json"

var (
	// Flags
	runMod        int
	PythonBinPath string
)

func main() {
	// parse all flags
	flag.IntVar(&runMod, "mod", 0, "0 to display stats about your provided movies json or 1 to run python scrtip to generate movies json file. Default: 0")
	flag.StringVar(&PythonBinPath, "py", "python", "Your local machine python bin path. Default python")
	flag.Parse()

	// OS env
	projectID := os.Getenv("GCP_PROJECT") // projectID := "golang-389808"
	dataSet := os.Getenv("GCP_DATASET")
	table := os.Getenv("GCP_TABLE")

	switch runMod {
	case 0:
		// parse movies json file & display stats about your movies list
		myMovies := parser.ParseJSON(filePath)
		internal.DisplayInfo(myMovies)
	case 1:
		parser.RunPyScript(PythonBinPath)
	case 3:
		internal.DoesFileExist(filePath)
	case 4:
		// Read from BigQuery
		sqltxt := internal.UserInput()

		if projectID == "" {
			log.Fatal("- GCP_PROJECT environment variable must be set.\n# bash: export GCP_PROJECT=xxx\n# powershell: $env:GCP_PROJECT=xxx")
		}
		c := bigquery.CreateBqClient(projectID)

		// bigquery.ReadFromBqDryRun(sqltxt, c)

		bigquery.ReadFromBq(sqltxt, c)
	case 5:
		// create dataset & table
		if projectID == "" || dataSet == "" || table == "" {
			log.Fatal("- GCP_PROJECT, GCP_DATASET, and GCP_TABLE environment variables must be set.\n# bash: export GCP_PROJECT=xxx\n# powershell: $env:GCP_PROJECT=xxx")
		}

		c := bigquery.CreateBqClient(projectID)
		bigquery.CreateDataSet(dataSet, c)
		bigquery.CreateTable(dataSet, table, c)

	case 6:
		// insert data
		if projectID == "" || dataSet == "" || table == "" {
			log.Fatal("- GCP_PROJECT, GCP_DATASET, and GCP_TABLE environment variables must be set.\n# bash: export GCP_PROJECT=xxx\n# powershell: $env:GCP_PROJECT=xxx")
		}

		c := bigquery.CreateBqClient(projectID)
		myMovies := parser.ParseJSON(filePath)
		bigquery.InsertIntoBq(dataSet, table, c, *myMovies)

	default:
		log.Fatal("invalid runMod")
	}
}
