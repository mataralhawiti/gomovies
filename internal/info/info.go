package info

import (
	"fmt"
	"runtime"
	"sort"

	"github.com/mataralhawiti/gomovies/internal/parser" // the directory Not the file!
)

// colors
var Yellow = "\033[33m"
var Reset = "\033[0m"

// DisplayInfo displays stats about your movies rating list
func DisplayInfo(moviesList *[]parser.Movie) {
	// ignore coloes if Windows
	if runtime.GOOS == "windows" {
		Yellow = ""
		Reset = ""
	}

	// overall movies rating
	moviesCount := parser.GetMoviesCount(moviesList)
	fmt.Printf("%v** You've rated (%v) movie, Great job!.%v\n", Yellow, moviesCount, Reset)

	// movies with given rating
	// moviesCountByRating, userInputR := parser.GetMoviesCountByRating(7, moviesList)
	// fmt.Printf("\nWe have (%v) movie out of (%v) with rating (%v).\n", moviesCountByRating, moviesCount, userInputR)

	// movies per rating, in descending order
	moviesRatingStats := parser.GetMoviesRatingStats(moviesList)
	fmt.Printf("%v** Movies count per rating as follow :%v\n", Yellow, Reset)

	keys := make([]int, 0, len(moviesRatingStats))
	for key := range moviesRatingStats {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {
		fmt.Printf(" - rating (%v) -> %v movies\n", k, moviesRatingStats[k])
	}

	// top 10 movies from the list by rating
	top10Movies := parser.GetTop10Movies(moviesList)
	fmt.Printf("%v** Your top 10 moives are :%v\n", Yellow, Reset)
	for i, m := range top10Movies {
		fmt.Printf(" %v. %v\n", i+1, m)
	}
}
