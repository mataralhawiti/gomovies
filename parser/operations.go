package parser

import "strconv"

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
func GetMoviesCountByRating(rating int, movieObj *[]Movie) int {
	sumMoive := 0

	for _, movie := range *movieObj {
		if movie.RATING == strconv.Itoa(rating) {
			sumMoive = sumMoive + 1
		}
	}
	return sumMoive
}

// MoviesCount
func GetMoviesCount(movieObj *[]Movie) int {
	return len(*movieObj)
}
