package adapter

import (
	"database/sql"
	"fmt"
)

type query struct {
	Table      string
	Parameters []any
	db         *sql.DB
}

func (q query) Insert() (string, error) {
	var insertedId string
	values := createValuesList(len(q.Parameters))

	err := q.db.QueryRow(
		fmt.Sprintf(
			"INSERT INTO %s VALUES (%s) RETURNING id",
			q.Table, values,
		),
		q.Parameters...,
	).Scan(&insertedId)

	if err != nil {
		return "", err
	}

	return insertedId, nil
}

func createValuesList(quantity int) string {
	var values string

	for i := 1; i <= quantity; i++ {
		values = values + fmt.Sprintf("$%d, ", i)
	}

	values = extractLastComma(values)
	return values
}

func extractLastComma(s string) string {
	return s[:len(s)-2]
}
