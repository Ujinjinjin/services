package factories

import (
	"fmt"
	"github.com/ujinjinjin/services/user/repository/context"
)

type DbContextFactory struct {
	connectionString string
}

func NewDbContextFactory(dbHost, dbUser, dbPassword, dbName *string) *DbContextFactory {
	return &DbContextFactory{
		connectionString: fmt.Sprintf("host=%s user=%s password=%s database=%s", *dbHost, *dbUser, *dbPassword, *dbName),
	}
}

// CreateUserDbContext creates database context
func (s *DbContextFactory) CreateUserDbContext() *context.UserDbContext {
	return context.NewUserDbContext(s.connectionString)
}
