package fruitshop

import (
	"context"
	user "fruitshop/frontend/gen/user"
	"log"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Add implements add.
func (s *usersrvc) Add(ctx context.Context, p *user.AddPayload) (res *user.UserManagement, err error) {
	res = &user.UserManagement{}
	s.logger.Print("user.add")
	return
}

// Get implements get.
func (s *usersrvc) Get(ctx context.Context, p *user.GetPayload) (res *user.UserManagement, err error) {
	res = &user.UserManagement{}
	s.logger.Print("user.get")
	return
}

// Show implements show.
func (s *usersrvc) Show(ctx context.Context) (res user.UserManagementCollection, err error) {
	s.logger.Print("user.show")
	return
}
