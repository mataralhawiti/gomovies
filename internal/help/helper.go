package help

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mataralhawiti/gomovies/internal/bigquery" // the directory Not the file!
)

// function to check if file exists
func DoesFileExist(fileName string) {
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
	return false, errors.New("invalid query number")
}

func UserInput(ds string, tb string) string {
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
				return fmt.Sprintf(sqltxt, ds, tb)
			}
		}
	} else {
		log.Fatal(err)
	}
	return sqltxt
}
