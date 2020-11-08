package factories

import "github.com/ujinjinjin/services/user/repository/context"

type DbContextFactory struct {

}

func NewDbContextFactory() *DbContextFactory {
	return &DbContextFactory{

	}
}

// CreateUser creates user with specified fields
func (s *DbContextFactory) CreateUserDbContext() (*context.UserDbContext, error) {
	return context.NewUserDbContext(), nil
}
