package internal

import (
	"fmt"
	"sort"

	"github.com/mataralhawiti/gomovies/parser" // the directory Not the file!
)

// DisplayInfo displays stats about your movies rating list
func DisplayInfo(moviesList *[]parser.Movie) {
	// overall movies rating
	moviesCount := parser.GetMoviesCount(moviesList)
	fmt.Printf("\nYou've rated (%v) movie, Great job!.\n", moviesCount)

	// movies with given rating
	moviesCountByRating, userInputR := parser.GetMoviesCountByRating(7, moviesList)
	fmt.Printf("\nWe have (%v) movie out of (%v) with rating (%v).\n", moviesCountByRating, moviesCount, userInputR)

	// movies per rating, in descending order
	moviesRatingStats := parser.GetMoviesRatingStats(moviesList)
	fmt.Println("\nMovies count per rating as follow :")

	keys := make([]string, 0, len(moviesRatingStats))
	for key := range moviesRatingStats {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, k := range keys {
		fmt.Printf("For rating %v ---> you have %v movie\n", k, moviesRatingStats[k])
	}

	// top 10 movies from the list by rating
	top10Movies := parser.GetTop10Movies(moviesList)
	fmt.Println(top10Movies)
}
