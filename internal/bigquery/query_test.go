// Fix this, test
package bigquery

import (
	"strconv"
)

var TestQList = map[string]string{
	"Movies count":                        "SELECT count(*) FROM `%s.%s`",
	"Movies count per year":               "SELECT count(*) FROM `%s.%s` WHERE year = \"2000\"",
	"Moives count grouped by year":        "SELECT year, count(movie_id) as ctn FROM `%s.%s` GROUP BY year ORDER BY count(movie_id) DESC",
	"display a movie info":                "SELECT name, year, rating, `desc` FROM `%s.%s` LIMIT 1",
	"the highest voting count":            "SELECT MAX(vote) FROM `%s.%s`",
	"Movie with the highest voting count": "TBI",
}

type TestQuery struct {
	id      string
	desc    string
	sqlText string
}

func TestQueriesList() []map[string]string {
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
		//QMap = map[string]string{}

		ctn++
	}

	return QSlice
}
