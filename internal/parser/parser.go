package parser

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Movie struct {
	MovieId   string
	MovieUrl  string
	MovieName string
	Year      int
	Cert      string
	Runtime   string
	Genre     []string
	Rating    int
	Desc      string
	Vote      int
}

func ParseJSON(filePath string) *[]Movie {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	//log.Printf("Succffully opend file %v\n", filePath)

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var movies []Movie
	// Unmarshal our JSON file
	json.Unmarshal(byteValue, &movies)

	return &movies
}
