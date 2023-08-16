package bigquery

var QList = map[string]string{
	"Count":             "SELECT count(movie_id) FROM movies.movies_full;",
	"CountPerYear":      "SELECT count(*) FROM movies.movies_full WHERE year = ?;",
	"CountGroupPerYear": "SELECT year, count(movie_id) as ctn FROM movies.movies_full GROUP BY year ORDER BY count(movie_id) DESC;",
	"display1":          "SELECT movie_id, movie_url, name, year, rating, `desc` FROM movies.movies_full LIMIT 1",
}

type Query struct {
	id      int
	desc    string
	sqlText string
}

// func (q *Query) QueriesList() []Query {
// 	QSlice := []Query{}
// 	ctn := 0

// 	for k, v := range QList {
// 		q.id = ctn
// 		q.desc = k
// 		q.sqlText = v

// 		QSlice = append(QSlice, *q)
// 	}

// 	return QSlice
// }

func QueriesList() []Query {
	var q Query
	QSlice := []Query{}
	ctn := 0

	for k, v := range QList {
		q.id = ctn
		q.desc = k
		q.sqlText = v

		QSlice = append(QSlice, q)
	}

	return QSlice
}
