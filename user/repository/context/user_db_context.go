package context

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/ujinjinjin/user_service/converters"
	"log"
)

type UserDbContext struct {
	connection *pgx.Conn
}

func NewUserDbContext() *UserDbContext {
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	var connection, err = pgx.Connect(context.Background(), "host=localhost user=local_user password=qwer1234 database=local_db")
	if err != nil {
		log.Printf("UserDbContext:NewUserDbContext:Error; %s", err)
	}

	return &UserDbContext{
		connection: connection,
	}
}

func (s *UserDbContext) Dispose() {
	err := s.connection.Close(context.Background())
	if err != nil {
		log.Printf("UserDbContext:Dispose:Error; %s", err)
	}
}

func (s *UserDbContext) TestSingle() (string, error) {
	testTable, err := converters.RowToTestTable(s.connection.QueryRow(context.Background(), "select id, name from test_table where id=1"))
	if err != nil {
		log.Printf("UserDbContext:TestSingle:Error; %v\n", err)
	}

	return testTable.Name, nil
}

func (s *UserDbContext) TestQuery() (string, error) {
	queryResult, err := s.connection.Query(context.Background(), "select id, name from test_table")
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
	}

	testTable, err := converters.RowsToTestTableArray(queryResult)
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
	}

	log.Printf("testTable: %v", testTable)

	return testTable[1].Name, nil
}