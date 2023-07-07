package parser

import (
	"sort"
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

	// sort map by desc before returning it
	keys := make([]string, 0, len(moviesStats))
	for k := range moviesStats {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	sortedMoviesStats := make(map[string]int16)
	for _, k := range keys {
		sortedMoviesStats[k] = moviesStats[k]
	}

	return sortedMoviesStats
}

func GetTop10Movies(movieObj *[]Movie) []string {
	/**
	1. detemine highest rating we have e.g. 10, 9, 8 etc
	2. get a list of movies with rating from step 1
	3. if list from step 2 is => 10, then get first 10 randomly and stop
	4. if not, detemine the reminaing number out of 10 and move to next highest rating.
	**/
	top10Movies := make([]string,0,10)

	moviebycnt := GetMoviesRatingStats(movieObj)

	keys := make([]string, 0, len(moviebycnt))
	for k := range moviebycnt {
		keys = append(keys, k)
	}
	//sort.Strings(keys)
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, r := range keys {
		intVar, _ := strconv.Atoi(r)
		tempSlice := GetMoviesByRating(intVar, movieObj)

		if len(top10Movies) < 10 {
			top10Movies = append(top10Movies, tempSlice...)
		} else {
			break
		}

	}

	return top10Movies[:10]
}
