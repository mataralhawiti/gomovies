package internal

import (
	"fmt"

	"github.com/mataralhawiti/gomovies/parser" // the directory Not the file!
)

// DisplayInfo displays stats about your movies rating list
func DisplayInfo(moviesList *[]parser.Movie) {
	moviesCount := parser.GetMoviesCount(moviesList)
	fmt.Printf("\nYou've rated (%v) movie, Great job!.\n",moviesCount)

	moviesCountByRating, userInputR := parser.GetMoviesCountByRating(7, moviesList)
	fmt.Printf("\nWe have (%v) movie out of (%v) with rating (%v).\n", moviesCountByRating, moviesCount, userInputR)

	moviesRatingStats := parser.GetMoviesRatingStats(moviesList)
	fmt.Println("\nMovies count per rating as follow :")
	for k,v := range moviesRatingStats {
		fmt.Printf("For rating %v ---> you have %v movie\n",k,v)
	}

	top10Movies := parser.GetTop10Movies(moviesList)
	fmt.Println(len(top10Movies))
}