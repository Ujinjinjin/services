package context

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/ujinjinjin/services/user/converters"
	"github.com/ujinjinjin/services/user/models"
	"log"
)

type UserDbContext struct {
	connection *pgx.Conn
}

// NewUserDbContext create new user database context
func NewUserDbContext(connectionString string) *UserDbContext {
	var connection, err = pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Printf("UserDbContext:NewUserDbContext:Error; %s", err)
	}

	return &UserDbContext{
		connection: connection,
	}
}

// Dispose close open connection to database
func (s *UserDbContext) Dispose() {
	err := s.connection.Close(context.Background())
	if err != nil {
		log.Printf("UserDbContext:Dispose:Error; %s", err)
	}
}

// GetUser get user
func (s *UserDbContext) GetUser(userId int32, isDeleted int32) (models.DbUser, error) {
	queryResult, err := s.connection.Query(
		context.Background(),
		"select * from user__get(p_user_id := $1, p_is_deleted := $2);",
		userId,
		isDeleted,
	)
	if err != nil {
		log.Printf("UserDbContext:GetUser:Error; %v\n", err)
		return models.DbUser{}, err
	}

	dbUserList, err := converters.RowsToDbUserArray(queryResult)
	if len(dbUserList) == 0 {
		return models.DbUser{}, nil
	}
	if err != nil {
		log.Printf("UserDbContext:GetUser:Error; %v\n", err)
		return models.DbUser{}, err
	}

	return dbUserList[0], nil
}

// GetUserList get user
func (s *UserDbContext) GetUserList(userIdList []int32, username string, email string, deleted int32) ([]models.DbUser, error) {
	queryResult, err := s.connection.Query(
		context.Background(),
		"select * from user__get_list(p_user_id_list := $1, p_username := $2, p_email := $3, p_is_deleted := $4);",
		userIdList,
		username,
		email,
		deleted,
	)
	if err != nil {
		log.Printf("UserDbContext:GetUser:Error; %v\n", err)
		return nil, err
	}

	dbUserList, err := converters.RowsToDbUserArray(queryResult)
	if err != nil {
		log.Printf("UserDbContext:GetUser:Error; %v\n", err)
		return nil, err
	}

	return dbUserList, nil
}
