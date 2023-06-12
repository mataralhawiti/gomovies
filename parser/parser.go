// https://stackoverflow.com/questions/53538123/how-to-parse-json-array-struct
package parser

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Movie struct {
	MOVIE_URL string
	Name      string
	YEAR      string
	RATING    string
	DESC      string
}

func ParseJSON(filePath string) *[]Movie {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our movies array
	var m []Movie

	// Unmarshal our JSON file
	json.Unmarshal(byteValue, &m)

	// print all
	//fmt.Println(m[0])

	// print Names
	// for _, data := range m {
	// 	fmt.Println(data.Name)
	// }

	return &m
}
