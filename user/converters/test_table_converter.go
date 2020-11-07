package converters

import (
	"github.com/jackc/pgx/v4"
	"github.com/ujinjinjin/user_service/models/db"
)

func RowToTestTable(row pgx.Row) (db.TestTable, error){
	var testTable db.TestTable
	err := row.Scan(&testTable.Id, &testTable.Name)
	if err != nil {
		return testTable, err
	}
	return testTable, nil
}

func RowsToTestTableArray(rows pgx.Rows) ([]db.TestTable, error){
	var result []db.TestTable

	for rows.Next() {
		var testTable db.TestTable
		var err = rows.Scan(&testTable.Id, &testTable.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, testTable)
	}

	return result, nil
}
