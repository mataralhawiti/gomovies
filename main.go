package main

import (
	"errors"
	"flag"
	"fmt"
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

	switch runMod {
	case 0:
		// parse movies json file & display stats about your movies list
		myMovies := parser.ParseJSON(filePath)
		internal.DisplayInfo(myMovies)
	case 1:
		parser.RunPyScript(PythonBinPath)
	case 3:
		doesFileExist(filePath)
	case 4:
		// Read from BigQuery
		sqltxt := userInput()

		projectID := "golang-389808"
		c := bigquery.CreateBqClient(projectID)
		bigquery.ReadFromBq(sqltxt, c)
	default:
		log.Fatal("invalid runMod")
	}
}

// function to check if file exists
func doesFileExist(fileName string) {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		fmt.Printf("%v file does not exist\n", fileName)
	} else {
		fmt.Printf("%v file exist\n", fileName)
	}
}

// check if value in Slice - Go 1.20
// For Go 1.21, we should use package:slices
func contains(slic []string, v string) (bool, error) {
	for _, s := range slic {
		if s == v {
			return true, nil
		}
	}
	return false, errors.New("Invalid query number")
}

func userInput() string {
	// BigQuery
	ql := bigquery.QueriesList()
	var ids []string

	// get QueriesList ids
	for _, k := range ql {
		ids = append(ids, k["id"])
	}

	// Take user input
	fmt.Println("Please enter query id .. ex: 2")
	for _, items := range ql {
		fmt.Printf("%v - %v\n", items["id"], items["desc"])
	}

	var qur string
	fmt.Scan(&qur)

	var sqltxt string

	// check if user's input is valid
	// if yes, extract SQL query text from predefined list
	if ok, err := contains(ids, qur); ok {
		for _, k := range ql {
			if sqltxt = k["sqlText"]; k["id"] == qur {
				return sqltxt
			}
		}
	} else {
		log.Fatal(err)
	}
	return sqltxt
}
