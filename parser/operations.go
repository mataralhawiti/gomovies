package parser

import (
	"strconv"
)

func GetMoiveNames(movieObj *[]Movie) []string {
	var names []string
	for _, movie := range *movieObj {
		names = append(names, movie.Name)
	}
	return names
}

// GetMovieByRating retrun a slice contains all movies names with the passed rating
func GetMoviesByRating(rating int, movieObj *[]Movie) []string {
	var namesByRating []string

	for _, movie := range *movieObj {
		if movie.RATING == strconv.Itoa(rating) {
			namesByRating = append(namesByRating, movie.Name)
		}
	}
	return namesByRating
}

// GetMovieByRating retrun count of moives with the given rating
func GetMoviesCountByRating(rating int, movieObj *[]Movie) (int, int) {
	cntMoivesByRating := 0

	for _, movie := range *movieObj {
		if movie.RATING == strconv.Itoa(rating) {
			cntMoivesByRating = cntMoivesByRating + 1
		}
	}
	return cntMoivesByRating, rating
}

// MoviesCount
func GetMoviesCount(movieObj *[]Movie) int {
	return len(*movieObj)
}

// GetMoviesRatingStats returns numbers of movies per rating
// rating range 1 to 10
func GetMoviesRatingStats(movieObj *[]Movie) map[string]int16 {
	moviesStats := make(map[string]int16)

	for _, movie := range *movieObj {
		moviesStats[movie.RATING] = moviesStats[movie.RATING] + 1
	}
	return moviesStats
}


func GetTop10Movies() map[string]int16{
	top10Movies := make(map[string]int16)

	return top10Movies
}