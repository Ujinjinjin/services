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
func (s *UserDbContext) GetUser(userId int32) (models.DbUser, error) {
	dbUser, err := converters.RowToDbUser(s.connection.QueryRow(context.Background(), "select * from user__get(p_user_id := $1);", userId))
	if err != nil {
		log.Printf("UserDbContext:GetUser:Error; %v\n", err)
		return models.DbUser{}, err
	}
	return dbUser, nil
}
