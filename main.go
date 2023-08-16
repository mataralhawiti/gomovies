package main

import (
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
		// BigQuery
		ql := bigquery.QueriesList()
		fmt.Println(ql)

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
