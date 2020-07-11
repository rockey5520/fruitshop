package users

import (
	"context"
	user "fruitshop/gen/user"
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
func (s *usersrvc) Add(ctx context.Context,
	p *user.AddPayload) (err error) {
	s.logger.Print("user.add started")
	newUser := user.UserManagement{
		MobieNumber: p.MobieNumber,
		UserName:    p.UserName,
	}
	err = CreateUser(&newUser)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("user.add completed")
	return
}

// Get implements get.
func (s *usersrvc) Get(ctx context.Context,
	p *user.GetPayload) (res *user.UserManagement, err error) {
	s.logger.Print("user.get started")
	result, err := GetUser(p.MobieNumber)
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("user.get completed")
	return &result, err
}

// Show implements show.
func (s *usersrvc) Show(ctx context.Context) (res user.UserManagementCollection,
	err error) {
	s.logger.Print("user.show started")
	res, err = ListUsers()
	if err != nil {
		s.logger.Print("An error occurred...")
		s.logger.Print(err)
		return
	}
	s.logger.Print("user.show completed")
	return
}
