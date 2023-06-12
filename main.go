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
}
