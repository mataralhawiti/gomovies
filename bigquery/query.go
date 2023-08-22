package bigquery

import (
	"strconv"
)

var QList = map[string]string{
	"Movies count":                 "SELECT count(movie_id) FROM {dataSet}.{table};",
	"Movies count per year":        "SELECT count(*) FROM movies.movies_full WHERE year = ?;",
	"Moives count grouped by year": "SELECT year, count(movie_id) as ctn FROM movies.movies_full GROUP BY year ORDER BY count(movie_id) DESC;",
	"display a movie info":         "SELECT name, year, rating, `desc` FROM movies.movies_full LIMIT 1",
	"display all movies data":      "SELECT name, year, rating, `desc` FROM movies.movies_full LIMIT 1",
}

type Query struct {
	id      string
	desc    string
	sqlText string
}

func QueriesList() []map[string]string {
	var q Query
	//QMap := map[string]string{}
	QSlice := []map[string]string{}
	ctn := 1

	for k, v := range QList {
		q.id = strconv.Itoa(ctn)
		q.desc = k
		q.sqlText = v
		QMap := map[string]string{"id": q.id, "desc": q.desc, "sqlText": q.sqlText}
		QSlice = append(QSlice, QMap)

		// clear QMap - running Go 1.20
		QMap = map[string]string{}

		ctn++
	}

	return QSlice
}
