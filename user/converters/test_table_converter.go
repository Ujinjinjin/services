package converters

import (
	"github.com/jackc/pgx/v4"
	"github.com/ujinjinjin/services/user/models"
)

func RowToTestTable(row pgx.Row) (models.TestTable, error){
	var testTable models.TestTable
	err := row.Scan(&testTable.Id, &testTable.Name)
	if err != nil {
		return testTable, err
	}
	return testTable, nil
}

func RowsToTestTableArray(rows pgx.Rows) ([]models.TestTable, error){
	var result []models.TestTable

	for rows.Next() {
		var testTable models.TestTable
		var err = rows.Scan(&testTable.Id, &testTable.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, testTable)
	}

	return result, nil
}
