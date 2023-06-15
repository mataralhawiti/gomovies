package main

import (
	"fmt"
	"os"

	"github.com/mataralhawiti/gomovies/parser" // the directory Not the file!
)

func main() {
	args := os.Args[1:]
	myMovies := parser.ParseJSON(args[0])

	//names := parser.GetMoiveNames(myMovies)

	// moviesByRating := parser.GetMoviesByRating(7, myMovies)
	// fmt.Println(moviesByRating)

	moviesCount := parser.GetMoviesCount(myMovies)
	fmt.Println(moviesCount)

	moviesCountByRating := parser.GetMoviesCountByRating(7, myMovies)
	fmt.Printf("We have %v movie out of %v with rating 7.\n", moviesCountByRating, moviesCount)

	moviesRatingStats := parser.GetMoviesRatingStats(myMovies)
	fmt.Println(moviesRatingStats)

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
