package converters

import (
	"github.com/jackc/pgx/v4"
	"github.com/ujinjinjin/services/user/models"
)

// RowToDbUser convert database query result to DbUser
func RowToDbUser(row pgx.Row) (models.DbUser, error){
	var dbUser models.DbUser
	err := row.Scan(&dbUser.UserId, &dbUser.Username, &dbUser.Email, &dbUser.FirstName, &dbUser.LastName, &dbUser.MiddleName)
	if err != nil {
		return dbUser, err
	}
	return dbUser, nil
}

// RowsToDbUserArray convert database query result to array of DbUser`s
func RowsToDbUserArray(rows pgx.Rows) ([]models.DbUser, error){
	var result []models.DbUser

	for rows.Next() {
		var dbUser models.DbUser
		var err = rows.Scan(&dbUser.UserId, &dbUser.Username, &dbUser.Email, &dbUser.FirstName, &dbUser.LastName, &dbUser.MiddleName)
		if err != nil {
			return nil, err
		}
		result = append(result, dbUser)
	}

	return result, nil
}
